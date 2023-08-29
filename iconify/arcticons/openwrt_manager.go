package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func OpenwrtManager(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M13.642 14.836c5.35-5.35 14.026-5.35 19.377 0c0 0 0 0 0 0M30.18 36.39a13.701 13.701 0 0 1-18.715-18.715m23.731-.001a13.701 13.701 0 0 1 .792 12.094m-17.875-10.46a7.378 7.378 0 0 1 10.434 0h0"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M9.17 10.365c7.82-7.82 20.5-7.82 28.32 0"/><circle cx="23.33" cy="24.525" r="2.635" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M41.457 38.13c.025-.193.039-.387.04-.582a4.807 4.807 0 0 0-.04-.58l1.258-.986a.29.29 0 0 0 .073-.38l-1.191-2.067a.29.29 0 0 0-.364-.128l-1.485.599a4.555 4.555 0 0 0-1.005-.581l-.224-1.578a.29.29 0 0 0-.29-.25h-2.39a.29.29 0 0 0-.29.25l-.224 1.578a4.47 4.47 0 0 0-1.005.58l-1.491-.598a.29.29 0 0 0-.363.128L31.274 35.6a.29.29 0 0 0 .073.38l1.255.986a4.848 4.848 0 0 0-.04.581c.001.195.015.389.04.581l-1.255.986a.29.29 0 0 0-.073.38l1.192 2.067a.29.29 0 0 0 .363.128l1.485-.6c.309.237.647.432 1.005.582l.224 1.578a.29.29 0 0 0 .29.25h2.384a.29.29 0 0 0 .29-.25l.224-1.578c.36-.149.697-.344 1.006-.581l1.485.599a.29.29 0 0 0 .363-.128l1.192-2.067a.29.29 0 0 0-.073-.38l-1.247-.986Zm-4.426 1.505a2.087 2.087 0 1 1 2.084-2.09v.003a2.084 2.084 0 0 1-2.08 2.087h-.004Z"/>`),
		g.Group(children),
	)
}