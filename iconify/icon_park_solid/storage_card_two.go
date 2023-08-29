package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StorageCardTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSStorageCardTwo0"><g fill="none" stroke-width="4"><rect width="32" height="40" x="8" y="4" stroke="#fff" rx="2"/><path stroke="#fff" stroke-linecap="round" stroke-linejoin="round" d="M14 4v10h20V4"/><path fill="#fff" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" d="M14 24h20v20H14z"/><path stroke="#000" stroke-linecap="round" stroke-linejoin="round" d="M14 32h20"/><path stroke="#fff" stroke-linecap="round" stroke-linejoin="round" d="M34 29v6m-20-6v6"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSStorageCardTwo0)"/>`),
		g.Group(children),
	)
}