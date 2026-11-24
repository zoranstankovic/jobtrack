# 05: Status Changes and timeline

**What to build:** On an Application, the user changes the Application Status with a small inline form (status + date, defaulting to today), both in the list and on the detail page, without a reload. Each change is added as a Status Change. The detail page shows the timeline, newest first. Any status may follow any other, and the same status may repeat (e.g. two Interview rounds). A Status Change dated in the past does not override a newer one when deciding the current Application Status. The user can delete a wrong Status Change, but not the last remaining one.

**Teaches:** append-only history tables, deriving current state from history (`ORDER BY` / `DISTINCT ON` or a lateral join), test-first development for ordering rules, updating several fragments from one HTMX response.

**Blocked by:** 04

**Status:** ready-for-agent

- [ ] The inline status form on list rows and on the detail page adds a Status Change and updates the shown current Application Status without a reload
- [ ] The detail page shows the full timeline (status, date), newest first
- [ ] Transitions are free: e.g. Ghosted → Interview and Planned → Rejected are accepted
- [ ] Adding a Status Change dated before the current latest one leaves the current Application Status unchanged
- [ ] On the same date, the Status Change recorded later counts as more recent
- [ ] Deleting a Status Change works, and deleting the only remaining one is refused with a message
- [ ] Tests written first (TDD) cover the current-status ordering rules, including past dates and same-date ties
