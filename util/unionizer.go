package util

import "reflect"

type Unionizer struct {
	items map[interface{}]bool
}

func (u *Unionizer) Len() int {
	return len(u.items)
}

func (u *Unionizer) JoinUnions(u2 Unionizer) Unionizer {
	out := Unionizer{}

	// add existing unions and return
	out.AddItems(u.GetUnion())
	out.AddItems(u2.GetUnion())

	return out
}

func (u *Unionizer) Contains(i interface{}) bool {
	_, ok := u.items[i]

	return ok
}

func (u *Unionizer) ForEach(each func(i interface{}) bool) {
	for k, _ := range u.items {
		if !each(k) {
			break
		}
	}
}

func (u *Unionizer) GetUnion() []interface{} {
	out := make([]interface{}, 0)

	for k := range u.items {
		out = append(out, k)
	}

	return out
}

// RemoveItems is a blanket removal of all items in the list.
func (u *Unionizer) RemoveItems(itemList interface{}) {
	rList := reflect.ValueOf(itemList)
	rLen := rList.Len()

	for i := 0; i < rLen; i++ {
		val := rList.Index(i).Interface()

		_, ok := u.items[val]

		if ok {
			delete(u.items, val)
		}
	}
}

// AddItems, on first run, adds the entire list to the union. On future runs, it subtracts all keys not present on the new list that were present on the current list of items
func (u *Unionizer) AddItems(itemList interface{}) {
	baseList := u.items == nil
	if baseList {
		u.items = make(map[interface{}]bool)
	}

	has := map[interface{}]bool{}

	rList := reflect.ValueOf(itemList)
	rLen := rList.Len()

	for i := 0; i < rLen; i++ {
		val := rList.Index(i).Interface()

		_, ok := u.items[val]

		if baseList || ok {
			has[val] = true
		}
	}

	if baseList {
		for k := range has {
			u.items[k] = true
		}
	} else {
		for k := range u.items {
			if _, ok := has[k]; !ok {
				delete(u.items, k)
			}
		}
	}
}
