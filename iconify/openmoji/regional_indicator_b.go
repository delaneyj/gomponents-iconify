package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RegionalIndicatorB(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<defs><path id="openmojiRegionalIndicatorB0" d="M46.826 29.022c0-4.398-3.578-7.976-7.976-7.976H28.174a1 1 0 0 0-1 1v27.908a1 1 0 0 0 1 1H38.85c4.398 0 7.976-3.578 7.976-7.976c0-3-1.666-5.618-4.122-6.978a7.978 7.978 0 0 0 4.122-6.978Zm-2 13.956a5.983 5.983 0 0 1-5.976 5.976h-9.676V37.006h9.676a5.98 5.98 0 0 1 5.976 5.972Zm-5.976-7.984c-.02 0-.038.01-.058.012h-9.618v-11.96h9.676a5.983 5.983 0 0 1 5.976 5.976a5.98 5.98 0 0 1-5.976 5.972Z"/></defs><use href="#openmojiRegionalIndicatorB0"/><g fill="none" stroke="#000" stroke-linejoin="round" stroke-width="2"><circle cx="36" cy="36" r="28"/><use href="#openmojiRegionalIndicatorB0"/></g>`),
		g.Group(children),
	)
}