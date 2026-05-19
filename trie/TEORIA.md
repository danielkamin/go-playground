# Trie (Prefix Tree / Drzewo prefiksowe)

## Czym jest Trie?

Trie to drzewiasta struktura danych do przechowywania ciągów znaków (stringów), zoptymalizowana pod kątem **operacji na prefiksach**. Nazwa pochodzi od słowa "re**trie**val".

Zamiast przechowywać każde słowo w całości (jak w hashsecie), Trie rozkłada je na litery i tworzy wspólne ścieżki dla słów o tym samym początku.

### Wizualizacja

Słowa: `["cat", "car", "card", "care", "bat"]`

```
        root
       /    \
      c      b
      |      |
      a      a
     / \     |
    t   r    t
        |\ 
        d  e
```

Każda ścieżka od `root` do węzła oznaczonego jako "koniec słowa" tworzy jedno słowo:
- `root → c → a → t` = "cat"
- `root → c → a → r → d` = "card"
- `root → c → a → r → e` = "care"

**Kluczowa obserwacja:** "cat" i "car" dzielą ścieżkę `c → a`. To oszczędza pamięć przy wielu słowach o wspólnym prefiksie.

---

## Struktura węzła

```go
type TrieNode struct {
    children [26]*TrieNode  // dla liter a-z
    isEnd    bool           // czy tu kończy się jakieś słowo
}
```

Alternatywnie zamiast tablicy `[26]` możesz użyć mapy `map[rune]*TrieNode` — bardziej ogólne, obsługuje unicode, ale wolniejsze.

### Dlaczego `[26]` a nie `map`?

| | `[26]*TrieNode` | `map[rune]*TrieNode` |
|--|--|--|
| Dostęp do dziecka | O(1), indeks = `ch - 'a'` | O(1) amortyzowane |
| Pamięć na węzeł | stała (26 × 8 bajtów) | dynamiczna |
| Obsługuje unicode | nie | tak |
| Sprawdzenie czy dziecko istnieje | `children[i] != nil` | `_, ok := children[ch]` |

Dla zadań z lowercase a-z → tablica. Dla ogólnego zastosowania → mapa.

---

## Operacje

### Insert(word string) — O(n), n = długość słowa

Iterujesz po literach słowa. Dla każdej litery:
- Jeśli dziecko dla tej litery nie istnieje → stwórz nowy węzeł
- Przejdź do dziecka
- Po ostatniej literze → ustaw `isEnd = true`

```
Insert("car"):
root → (brak 'c'? utwórz) → c → (brak 'a'? utwórz) → a → (brak 'r'? utwórz) → r
                                                                                   isEnd = true
```

### Search(word string) bool — O(n)

Iterujesz po literach. Jeśli w którymkolwiek kroku dziecko nie istnieje → `false`. Po ostatniej literze sprawdzasz `isEnd`.

**Uwaga na pułapkę:** "car" i "card" są w Trie. Szukasz "car" — dojdziesz do węzła `r`, ale czy `isEnd == true`? To zależy od tego czy "car" było wstawione. Samo dotarcie do węzła nie wystarczy.

### StartsWith(prefix string) bool — O(n)

Identycznie jak `Search`, ale na końcu **nie** sprawdzasz `isEnd` — interesuje cię tylko czy istnieje ścieżka dla prefiksu.

### GetWordsWithPrefix(prefix string) []string — O(n + k)

To serce autocomplete. Algorytm:
1. Znajdź węzeł odpowiadający ostatniej literze prefiksu (jak `StartsWith`)
2. Od tego węzła wykonaj **DFS** (depth-first search) — zbierz wszystkie słowa kończące się `isEnd = true`

---

## Złożoność

| Operacja | Czas | Pamięć |
|----------|------|--------|
| Insert | O(n) | O(n) na nowe węzły |
| Search | O(n) | O(1) |
| StartsWith | O(n) | O(1) |
| GetWordsWithPrefix | O(p + k) | O(k) na wyniki |

`n` = długość słowa/prefiksu, `p` = długość prefiksu, `k` = liczba znaków we wszystkich pasujących słowach

**Porównanie z hashmapą:**
- `map[string]bool` → `Search` w O(1), ale `StartsWith` wymaga iteracji po wszystkich kluczach O(n×m)
- Trie → `Search` w O(długość_słowa), ale `StartsWith` bardzo efektywne

---

## Kiedy używać?

- **Autocomplete** (wyszukiwarki, IDE, terminale)
- **Spell checker** — sprawdzanie czy słowo istnieje w słowniku
- **IP routing** — routery używają Trie na bitach adresu IP (longest prefix match)
- **Word games** — Boggle, Scrabble (sprawdzenie czy prefiks może być początkiem słowa)

---

## Zadanie: Zaimplementuj Trie od zera

### Interfejs do zaimplementowania

```go
type Trie struct { ... }

func NewTrie() *Trie
func (t *Trie) Insert(word string)
func (t *Trie) Search(word string) bool
func (t *Trie) StartsWith(prefix string) bool
func (t *Trie) GetWordsWithPrefix(prefix string) []string  // dla autocomplete
```

### Kroki implementacji

**Krok 1 — Węzeł:**
```go
type TrieNode struct {
    children [26]*TrieNode
    isEnd    bool
}
```
Litera `ch` mapuje się na indeks przez: `ch - 'a'` (zakładamy lowercase ASCII).

**Krok 2 — Insert:**
- Zacznij od `root`
- Dla każdej litery `ch` w `word`:
  - Oblicz indeks: `idx := ch - 'a'`
  - Jeśli `current.children[idx] == nil` → utwórz nowy `TrieNode`
  - Przejdź: `current = current.children[idx]`
- Na końcu: `current.isEnd = true`

**Krok 3 — Search:**
- Tak samo jak Insert, ale bez tworzenia węzłów
- Jeśli `children[idx] == nil` → zwróć `false`
- Na końcu zwróć `current.isEnd`

**Krok 4 — StartsWith:**
- Identycznie jak Search, ale na końcu zwróć `true` (nie sprawdzaj `isEnd`)

**Krok 5 — GetWordsWithPrefix (trudniejszy):**
- Znajdź węzeł odpowiadający ostatniej literze prefiksu (jak StartsWith, ale zachowaj wskaźnik na węzeł)
- Napisz pomocniczą funkcję `dfs(node *TrieNode, current string, results *[]string)`:
  - Jeśli `node.isEnd` → dodaj `current` do `results`
  - Dla każdego `i` od 0 do 25: jeśli `node.children[i] != nil` → rekurencja z `current + string('a'+i)`

### Hint — nawigacja do węzła prefiksu

```go
func (t *Trie) findNode(prefix string) *TrieNode {
    current := t.root
    for _, ch := range prefix {
        idx := ch - 'a'
        if current.children[idx] == nil {
            return nil  // prefiks nie istnieje
        }
        current = current.children[idx]
    }
    return current
}
```
Użyj tej pomocniczej funkcji zarówno w `StartsWith` jak i `GetWordsWithPrefix`.

### Testy do napisania

```go
func TestTrieInsertSearch(t *testing.T)       // podstawowe insert + search
func TestTrieSearchMissing(t *testing.T)      // szukanie nieistniejącego słowa
func TestTriePrefix(t *testing.T)             // StartsWith dla istniejących i brakujących prefiksów
func TestTriePrefixVsWord(t *testing.T)       // "car" jako prefiks ale nie słowo (tylko "card" wstawione)
func TestTrieAutocomplete(t *testing.T)       // GetWordsWithPrefix zwraca wszystkie pasujące słowa
func TestTrieEmptyPrefix(t *testing.T)        // GetWordsWithPrefix("") zwraca wszystkie słowa
```

### Czego NIE wolno użyć
- `strings.HasPrefix` — sens ćwiczenia to własna struktura
- Żadnej gotowej implementacji drzewa

### Rozszerzenie (opcjonalne)

Po zrobieniu podstawowej wersji możesz spróbować:
1. **Delete(word)** — usunięcie słowa z Trie (trudniejsze — trzeba uważać żeby nie usunąć wspólnych prefiksów)
2. **CountWordsWithPrefix(prefix)** — zlicz słowa bez zbierania ich wszystkich (przydaj pole `count` w węźle)
3. **Concurrent-safe Trie** — tak samo jak LRU cache, dodaj `sync.RWMutex`
