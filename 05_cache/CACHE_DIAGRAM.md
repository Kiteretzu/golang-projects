# LRU Cache System — How It Works

This cache is a **Least Recently Used (LRU)** cache with a fixed size (`SIZE = 5`). It uses a **doubly linked list** for order (recency) and a **hash map** for O(1) lookup.

---

## 1. Structure Overview

```
┌─────────────────────────────────────────────────────────────────┐
│                          CACHE                                    │
│  ┌───────────────────────────────────┐  ┌─────────────────────┐ │
│  │  QUEUE (doubly linked list)        │  │  HASH                │ │
│  │  Order = recency (front = recent)   │  │  string → *Node      │ │
│  │                                    │  │  O(1) lookup         │ │
│  │  Head ⇄ [node] ⇄ ... ⇄ [node] ⇄ Tail│  └─────────────────────┘ │
│  │         ↑ most recent    least ↑   │                           │
│  └───────────────────────────────────┘                           │
└─────────────────────────────────────────────────────────────────┘
```

- **Node**: `Val` (string), `Left`, `Right` (pointers).
- **Queue**: `Head` and `Tail` are dummy nodes; real items are between them. `Head.Right` = most recent, `Tail.Left` = least recent.
- **Hash**: Maps each cached string to its `*Node` in the list.

---

## 2. Check(str) — Main Entry Point

```
                    ┌──────────────┐
                    │  Check(str)  │
                    └──────┬───────┘
                           │
                    ┌──────▼──────┐
                    │ str in Hash?│
                    └──────┬──────┘
              ┌────────────┼────────────┐
              │ Yes                    No
              ▼                         ▼
    ┌─────────────────┐      ┌─────────────────┐
    │ Remove(node)     │      │ Create Node(str)│
    │ from list & Hash │      │                 │
    └────────┬─────────┘      └────────┬─────────┘
             │                        │
             └────────────┬───────────┘
                          ▼
                 ┌─────────────────┐
                 │ Add(node) at     │
                 │ front (Head.Right)│
                 └────────┬─────────┘
                          │
                 ┌────────▼────────┐
                 │ Hash[str] = node│
                 └────────┬────────┘
                          │
                 ┌────────▼────────┐
                 │ Length > SIZE?   │
                 │ → Remove(Tail.Left)│
                 │   (evict LRU)   │
                 └─────────────────┘
```

- **If `str` is in the cache**: its node is **removed** from the list (and from Hash in `Remove`), then **re-added** at the front so it becomes “most recently used”.
- **If `str` is not in the cache**: a **new node** is created, then **added** at the front and stored in **Hash**.
- **Add** always inserts at the front. If the list length exceeds `SIZE`, **Remove(Tail.Left)** is called to evict the least recently used item (and it’s removed from Hash inside `Remove`).

---

## 3. Queue Layout (Doubly Linked List)

```
     Head                    Tail
       │                      │
       ▼                      ▼
   [dummy] ⇄ [parrot] ⇄ [avocado] ⇄ [dragonfruit] ⇄ [tree] ⇄ [potato] ⇄ [dummy]
              ↑ most recent              least recent ↑
```

- New or re-used items are inserted at **Head.Right**.
- When full, the node at **Tail.Left** is removed (eviction).

---

## 4. Remove(node)

- Unlinks `node` from the list: `node.Left.Right = node.Right`, `node.Right.Left = node.Left`.
- Decrements `Queue.Length`.
- `delete(Hash, node.Val)`.
- Returns the unlinked node (so it can be re-used in **Add**).

---

## 5. Add(node)

- Inserts `node` at the front: between `Head` and `Head.Right`.
- Increments `Queue.Length`.
- If `Queue.Length > SIZE`, calls **Remove(Queue.Tail.Left)** to evict the LRU item.

---

## 6. Example Run (from main)

Words processed in order: `parrot`, `avocado`, `dragonfruit`, `tree`, `potato`, `tomato`, `tree`, `dog`.

| Step | Word        | Action                          | Queue (front → back)                    |
|------|-------------|----------------------------------|-----------------------------------------|
| 1    | parrot      | Add                             | [parrot]                                 |
| 2    | avocado     | Add                             | [avocado, parrot]                        |
| 3    | dragonfruit | Add                             | [dragonfruit, avocado, parrot]          |
| 4    | tree        | Add                             | [tree, dragonfruit, avocado, parrot]    |
| 5    | potato      | Add                             | [potato, tree, dragonfruit, avocado, parrot] |
| 6    | tomato      | Add, evict parrot               | [tomato, potato, tree, dragonfruit, avocado] |
| 7    | tree        | Hit — move to front             | [tree, tomato, potato, dragonfruit, avocado] |
| 8    | dog         | Add, evict avocado              | [dog, tree, tomato, potato, dragonfruit] |

So the **caching system** works by: (1) **Check** updates recency or inserts, (2) **Queue** keeps order (front = MRU, back = LRU), (3) **Hash** gives O(1) access to the node, (4) when over capacity, the node at **Tail.Left** is **Removed** (evicted).
