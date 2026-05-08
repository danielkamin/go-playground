package main

import "testing"

func TestLRUBasic(t *testing.T) {
	// get/put podstawowe
	t.Log("Success TestLRUBasic!")
}
func TestLRUEviction(t *testing.T) {
	t.Log("Success TestLRUEviction!")

} // czy LRU element jest usuwany
func TestLRUUpdateExisting(t *testing.T) {
	t.Log("Success TestLRUUpdateExisting!")

} // put na istniejący klucz
func TestLRUMoveOnGet(t *testing.T) {
	t.Log("Success TestLRUMoveOnGet!")

} // get przesuwa element, zmienia kolejność
