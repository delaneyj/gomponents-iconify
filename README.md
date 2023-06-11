# All the icons in a single place

<img alt="mascot" src="assets/mascot.png" width="384px" />

[![Go Reference](https://pkg.go.dev/badge/github.com/delaneyj/gomponents-iconify.svg)](https://pkg.go.dev/github.com/delaneyj/gomponents-iconify)

# gomponents-iconify
Iconify as Gomponents.  Similar to Maragudk's [gomponents-heroicons](https://github.com/maragudk/gomponents-heroicons) but is a superset of all the icons available from [Iconify](https://iconify.design/).   Currently the best way to search for icons is using https://icones.js.org/.

For example [Material Design's Logout](https://icones.js.org/collection/all?s=mdi:logout) is available as `mdi.Logout` from `github.com/delaneyj/gomponents-iconify/mdi`.

Also each iconset has a `IconFromName` function that will return a `gomponents.Node` for the icon.  For example `mdi.IconFromName("logout")` will return the same `gomponents.Element` as `mdi.Logout`.  Useful for dynamic icon selection.