package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Wechat(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><path d="M36.997 21.711C36.843 13.008 29.74 6 21 6C12.163 6 5 13.163 5 22c0 4.17 1.595 7.968 4.209 10.815l-1.199 7.21l7.115-3.055c3.135 1.042 6.093 1.303 8.875.782" clip-rule="evenodd"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M36.997 21.711C36.843 13.008 29.74 6 21 6C12.163 6 5 13.163 5 22c0 4.17 1.595 7.968 4.209 10.815l-1.199 7.21l7.115-3.055c3.135 1.042 6.093 1.303 8.875.782"/><path fill="currentColor" d="M15.125 20.467a2.258 2.258 0 0 0 2.25-2.267a2.258 2.258 0 0 0-2.25-2.267a2.258 2.258 0 0 0-2.25 2.267a2.258 2.258 0 0 0 2.25 2.267Zm9 0a2.258 2.258 0 0 0 2.25-2.267a2.258 2.258 0 0 0-2.25-2.267a2.258 2.258 0 0 0-2.25 2.267a2.258 2.258 0 0 0 2.25 2.267Z"/><path d="M38.762 39.93A10.453 10.453 0 0 1 32.5 42C26.701 42 22 37.299 22 31.5S26.701 21 32.5 21S43 25.701 43 31.5c0 1.6-.358 3.116-.998 4.473" clip-rule="evenodd"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M38.762 39.93A10.453 10.453 0 0 1 32.5 42C26.701 42 22 37.299 22 31.5S26.701 21 32.5 21S43 25.701 43 31.5c0 1.6-.358 3.116-.998 4.473"/><path d="M42.002 35.973L43 42l-4.238-2.07" clip-rule="evenodd"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M42.002 35.973L43 42l-4.238-2.07"/><path fill="currentColor" d="M35.688 30.8A1.694 1.694 0 0 1 34 29.1c0-.939.755-1.7 1.688-1.7c.931 0 1.687.761 1.687 1.7s-.755 1.7-1.688 1.7Zm-6.75 0a1.694 1.694 0 0 1-1.688-1.7c0-.939.756-1.7 1.688-1.7c.931 0 1.687.761 1.687 1.7s-.756 1.7-1.688 1.7Z"/></g>`),
		g.Group(children),
	)
}