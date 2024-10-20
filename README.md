# All the icons in a single place

<img alt="mascot" src="assets/mascot.png" width="384px" />

## Warning

This generates a LOT of code, first time you use a new iconset it will take a while to compile.  Subsequent compiles will be much faster but be warned.

[![Go Reference](https://pkg.go.dev/badge/github.com/delaneyj/gomponents-iconify.svg)](https://pkg.go.dev/github.com/delaneyj/gomponents-iconify)

# gomponents-iconify
Iconify as Gomponents.  Similar to Maragudk's [gomponents-heroicons](https://maragu.dev/gomponents-heroicons) but is a superset of all the icons available from [Iconify](https://iconify.design/).   Currently the best way to search for icons is using https://icones.js.org/.

For example [Material Design's Logout](https://icones.js.org/collection/all?s=mdi:logout) is available as `mdi.Logout` from `github.com/delaneyj/gomponents-iconify/mdi`.

Also each iconset has a `IconFromName` function that will return a `gomponents.Node` for the icon.  For example `mdi.IconFromName("logout")` will return the same `gomponents.Element` as `mdi.Logout`.  Useful for dynamic icon selection.