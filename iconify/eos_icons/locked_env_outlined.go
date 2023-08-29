package eos_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LockedEnvOutlined(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7.829 14.59v.752H6.524v.877H8V17H5.55v-4H8v.781H6.524v.809h1.305z"/><path fill="currentColor" d="M17 8h-1V6A5 5 0 0 0 6 6v2H5a2.006 2.006 0 0 0-2 2v10a2.006 2.006 0 0 0 2 2h12a2.006 2.006 0 0 0 2-2V10a2.006 2.006 0 0 0-2-2ZM7.9 6a3.1 3.1 0 1 1 6.2 0v2H7.9ZM17 20H5V10h12Z"/><path fill="currentColor" d="M12.168 17h-.974l-1.63-2.467V17H8.59v-4h.974l1.63 2.479V13h.974v4zm2.933 0h-1.202L12.5 13h1l1 3l1-3h1l-1.399 4z"/>`),
		g.Group(children),
	)
}