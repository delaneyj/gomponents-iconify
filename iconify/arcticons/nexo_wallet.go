package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NexoWallet(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M14.11 18.29L43.5 35.259l-9.826 5.672l-9.976-5.76l.02-.011l-9.608-5.546Zm0 11.324v.022l9.588 5.536l-9.608 5.504v.015l-.013-.008l.013-.007l.01-11.068Zm-.047 11.062v.015l.014-.008l-.014-.007l.037-11.068l.01-11.318l-9.61-5.549v22.413ZM33.937 7.324V7.31l-.015.008l.015.007l-.037 11.068l-.01 11.318l9.61 5.548V12.846Zm-9.635 5.505l-.02.01l9.608 5.547V29.71L4.5 12.741l9.826-5.672Zm0 0l9.607-5.505V7.31l.013.008l-.013.007l-.009 11.068l-.01-.006v-.022Z"/>`),
		g.Group(children),
	)
}