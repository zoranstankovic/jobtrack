# 09: Refresh one Search Profile from BA

**What to build:** On a Search Profile, the user clicks Refresh. jobtrack asks the Bundesagentur für Arbeit for matching Job Postings (at most 2 pages of 50) and stores each new one once, identified by (Source, `referenznummer`). It records that this Search Profile found it and shows "N new postings". The profile page lists the Job Postings it found: title, employer name, location, publication date, and a link to the original posting on the BA website. Refreshing again does not create duplicates and does not change existing postings. If BA is unreachable or returns an error or invalid data, the user sees a clear message and nothing already stored is lost.

The BA Source client is the only code that knows the BA API. It translates a Search Profile into query parameters: keywords → `was`, location → `wo`, radius → `umkreis`, working time → `arbeitszeit` (`vz`, `tz`, `ho` joined by `;`), days back → `veroeffentlichtseit`, and always `angebotsart=1`. It sends the `X-API-Key: jobboerse-jobsuche` header to `GET /pc/v6/jobs` and reads `ergebnisliste` (`referenznummer`, `stellenangebotsTitel`, `firma`, `arbeitgeberKundennummerHash`, `stellenlokationen[].adresse`, `datumErsteVeroeffentlichung`). See the spec for the verified API facts.

**Teaches:** HTTP clients in Go (timeouts, headers, query building), JSON decoding into structs, pagination, `httptest` fake servers with fixtures recorded from real responses, `INSERT … ON CONFLICT` upserts, test-first development.

**Blocked by:** 08

**Status:** ready-for-agent

- [ ] Written test-first (TDD) against the fake BA server through the HTTP seam
- [ ] The requests the fake server receives contain the correctly translated query parameters and the API key header (including several working times joined with `;` and no `arbeitszeit` when none are selected)
- [ ] At most 2 pages of 50 are requested; fetching stops early when the results run out
- [ ] New Job Postings are stored with title, employer name, BA employer identifier, location, publication date and Posting Status New
- [ ] A posting that already exists is not duplicated and its fields and Posting Status are not changed; it is linked to this Search Profile if it wasn't yet
- [ ] The response shows the number of new postings; the profile page lists its Job Postings with a link to the original on the BA website
- [ ] A BA error (5xx, timeout, invalid JSON) shows a clear message and leaves stored data untouched
- [ ] Fixtures are taken from real BA responses
