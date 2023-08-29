package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UsbSquareBroken(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none"><circle cx="12" cy="17" r="1" stroke="currentColor" stroke-width="1.5"/><circle cx="8" cy="9" r="1" stroke="currentColor" stroke-width="1.5"/><path fill="currentColor" d="M11.25 16a.75.75 0 0 0 1.5 0h-1.5ZM12 6l.53-.53a.75.75 0 0 0-1.06 0L12 6Zm.47 1.53a.75.75 0 1 0 1.06-1.06l-1.06 1.06Zm-2-1.06a.75.75 0 0 0 1.06 1.06l-1.06-1.06Zm1 .06l1 1l1.06-1.06l-1-1l-1.06 1.06Zm0-1.06l-1 1l1.06 1.06l1-1l-1.06-1.06ZM12.75 16v-1.875h-1.5V16h1.5Zm0-1.875V6h-1.5v8.125h1.5Z"/><path stroke="currentColor" stroke-linecap="round" stroke-width="1.5" d="M8 10v1.03a2 2 0 0 0 1.403 1.909l1.404.438A1.7 1.7 0 0 1 12 15m4-4v1.03a2 2 0 0 1-1.403 1.909l-1.404.438A1.7 1.7 0 0 0 12 16"/><path stroke="currentColor" stroke-width="1.5" d="M15 10c0-.471 0-.707.146-.854C15.293 9 15.53 9 16 9s.707 0 .854.146C17 9.293 17 9.53 17 10s0 .707-.146.854C16.707 11 16.47 11 16 11s-.707 0-.854-.146C15 10.707 15 10.47 15 10Z"/><path stroke="currentColor" stroke-linecap="round" stroke-width="1.5" d="M22 12c0 4.714 0 7.071-1.465 8.535C19.072 22 16.714 22 12 22s-7.071 0-8.536-1.465C2 19.072 2 16.714 2 12s0-7.071 1.464-8.536C4.93 2 7.286 2 12 2c4.714 0 7.071 0 8.535 1.464c.974.974 1.3 2.343 1.41 4.536"/></g>`),
		g.Group(children),
	)
}