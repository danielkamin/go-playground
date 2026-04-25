# Hashmap, LRU Cache, Doubly Linked List

## 1. Hash Map / Hash Table

Struktura danych przechowująca pary klucz-wartość z dostępem w czasie O(1) amortyzowanym.

**Jak działa:**
1. Masz tablicę "kubełków" (buckets) o stałym rozmiarze (np. 16)
2. Dla każdego klucza wywołujesz **funkcję hashującą** → indeks (`hash(key) % len(buckets)`)
3. Pod tym indeksem przechowujesz parę `{key, value}`
4. **Kolizja** = dwa klucze trafiają na ten sam indeks

**Rozwiązania kolizji:**
- **Chaining** — każdy bucket to lista (linked list) par klucz-wartość
- **Open addressing** — przy kolizji szukasz kolejnego wolnego slotu w tablicy

**Złożoność:**
| Operacja | Średnia | Worst case |
|----------|---------|------------|
| Get      | O(1)    | O(n)       |
| Put      | O(1)    | O(n)       |
| Delete   | O(1)    | O(n)       |

Worst case zachodzi gdy wszystkie klucze trafią do jednego bucketa (np. zła funkcja hashująca).

**Load factor** = `liczba_elementów / liczba_bucketów`. Gdy przekroczy próg (np. 0.75), hashmap rozrasta się (rehashing) — tworzy nową, większą tablicę i przenosi wszystkie elementy.

---

## 2. Doubly Linked List (Lista dwukierunkowa)

Lista, w której każdy węzeł zna zarówno **następnika** jak i **poprzednika**.

```
nil ← [prev|A|next] ↔ [prev|B|next] ↔ [prev|C|next] → nil
       head                                  tail
```

**Węzeł:**
```go
type Node struct {
    key, value string
    prev, next *Node
}
```

**Zaleta nad listą jednokierunkową:** usunięcie węzła ze środka listy w O(1) — nie trzeba iterować od początku, bo węzeł zna swojego poprzednika.

**Operacje:**
- `AddToFront(node)` — O(1)
- `Remove(node)` — O(1) — kluczowe dla LRU
- `RemoveLast()` — O(1) — kluczowe dla LRU

---

## 3. LRU Cache (Least Recently Used)

Cache o ograniczonej pojemności, który przy przepełnieniu usuwa **najdawniej używany** element.

**Zasada działania:**
- Każde `Get` i `Put` przesuwa element na pozycję "najnowszego"
- Gdy cache jest pełny i dodajesz nowy element → usuwasz "najstarszy" (ten na końcu listy)

**Dlaczego to realne?**
- Redis używa LRU jako strategię eksmisji kluczy
- CPU L1/L2 cache działa na podobnej zasadzie
- Przeglądarki cachują zasoby HTTP w ten sposób

**Implementacja = HashMap + Doubly Linked List:**
```
HashMap:  klucz → wskaźnik na węzeł listy   (O(1) dostęp)
DLL:      kolejność użycia, head=najnowszy, tail=najstarszy
```

```
Put("a", "1"):  HashMap["a"] → Node{a,1}   DLL: [a]
Put("b", "2"):  HashMap["b"] → Node{b,2}   DLL: [b, a]
Put("c", "3"):  HashMap["c"] → Node{c,3}   DLL: [c, b, a]
Get("a"):       przesuń a na przód          DLL: [a, c, b]
Put("d", "4"):  cache pełny (cap=3)         DLL: [d, a, c]  ← b usunięty
```

---

## Zadanie: Zaimplementuj LRU Cache od zera

**Cel:** Zbudować LRU Cache bez użycia wbudowanego `map` z Go — własny hashmap + własna doubly linked list.

### Interfejs do zaimplementowania

```go
type LRUCache struct { ... }

func NewLRUCache(capacity int) *LRUCache
func (c *LRUCache) Get(key string) (string, bool)
func (c *LRUCache) Put(key, value string)
```

### Kroki implementacji

**Krok 1 — Doubly Linked List**
- Struct `Node` z polami: `key`, `value`, `prev`, `next`
- Funkcje: `addToFront(node)`, `remove(node)`, `removeLast() *Node`
- Użyj dummy `head` i `tail` (sentinel nodes) — upraszcza edge case'y

**Krok 2 — Własny HashMap**
- Struct `HashMap` z tablicą bucketów (`[]*Entry`) i rozmiarem
- Każdy bucket to linked list par `{key, value, next}`
- Funkcje: `get(key)`, `set(key, node)`, `delete(key)`
- Funkcja hashująca: możesz użyć prostej sumy bajtów klucza `% len(buckets)`

**Krok 3 — LRU Cache**
- Połącz HashMap (przechowuje `key → *Node`) i DLL (kolejność użycia)
- `Get(key)`:
  1. Znajdź węzeł przez HashMap
  2. Przesuń go na front DLL
  3. Zwróć wartość
- `Put(key, value)`:
  1. Jeśli klucz istnieje → zaktualizuj wartość, przesuń na front
  2. Jeśli nie istnieje → stwórz węzeł, dodaj do frontu DLL i do HashMap
  3. Jeśli `len > capacity` → usuń ostatni węzeł z DLL i z HashMap

### Testy do napisania

```go
func TestLRUBasic(t *testing.T)        // get/put podstawowe
func TestLRUEviction(t *testing.T)     // czy LRU element jest usuwany
func TestLRUUpdateExisting(t *testing.T) // put na istniejący klucz
func TestLRUMoveOnGet(t *testing.T)    // get przesuwa element, zmienia kolejność eksmisji
```

### Czego NIE wolno użyć
- `map` z Go
- `container/list` z biblioteki standardowej

### Hint — sentinel nodes w DLL

Zamiast sprawdzać `if head == nil` wszędzie, użyj dummy węzłów:
```
dummyHead ↔ [rzeczywiste węzły] ↔ dummyTail
```
Wtedy `addToFront` zawsze wstawia za `dummyHead`, a `removeLast` zawsze usuwa węzeł przed `dummyTail`.
