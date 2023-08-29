package cil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Tag(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M485.887 263.261L248 25.373A31.791 31.791 0 0 0 225.373 16H64a48.055 48.055 0 0 0-48 48v161.078A32.115 32.115 0 0 0 26.091 248.4l253.061 237.725a23.815 23.815 0 0 0 16.41 6.51q.447 0 .9-.017a23.828 23.828 0 0 0 16.79-7.734l173.329-188.405a23.941 23.941 0 0 0-.694-33.218ZM295.171 457.269L48 225.078V64a16.019 16.019 0 0 1 16-16h161.373l232.461 232.462Z"/><path fill="currentColor" d="M148 96a52 52 0 1 0 52 52a52.059 52.059 0 0 0-52-52Zm0 72a20 20 0 1 1 20-20a20.023 20.023 0 0 1-20 20Z"/>`),
		g.Group(children),
	)
}