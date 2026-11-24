# 06: Edit/delete Applications, active filter, Company page

**What to build:** The user edits all fields of an Application and deletes an Application after confirming. The Applications list has an "Active only" filter that hides Accepted, Rejected, Withdrawn and Ghosted. Each Company has its own page with its details and all its Applications (several per Company are allowed). The Companies list shows the number of Applications. A Company that still has Applications cannot be deleted, and the user is told why.

**Teaches:** filtering with query parameters and SQL `WHERE`, `JOIN` and `GROUP BY` / `COUNT`, handling foreign key restrictions as user-facing errors, `hx-push-url` for filter state.

**Blocked by:** 03, 04

**Status:** ready-for-agent

- [ ] The Application edit form updates every field from 04 with the same validation; the Status Changes are not affected
- [ ] Deleting an Application asks for confirmation and removes it together with its Status Changes
- [ ] The "Active only" filter on the Applications list excludes Accepted, Rejected, Withdrawn and Ghosted, and the filter state is kept in the URL
- [ ] The Company page lists all its Applications with current Application Status
- [ ] The Companies list shows each Company's number of Applications
- [ ] Deleting a Company that has Applications is refused with a clear message; a Company without Applications can still be deleted
- [ ] Tests through the HTTP seam cover: edit, delete, the active filter, refused Company deletion
