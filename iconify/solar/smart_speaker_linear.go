package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SmartSpeakerLinear(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none"><path fill="currentColor" d="M2.507 14.615c.877.765 3.318 2.635 9.493 2.635v-1.5c-5.807 0-7.9-1.736-8.507-2.266l-.986 1.131ZM12 17.25c3.15 0 5.325-.48 6.823-1.066c1.509-.59 2.3-1.278 2.684-1.632l-1.015-1.104c-.297.273-.932.836-2.216 1.339c-1.294.507-3.278.963-6.276.963v1.5Z"/><path stroke="currentColor" stroke-width="1.5" d="M15.5 16c-.582-1.748-1.653-2.5-3.5-2.5s-2.918.752-3.5 2.5"/><path stroke="currentColor" stroke-linecap="round" stroke-width="1.5" d="M8 4.412c-.32.275-.5.673-.5 1c0 1.153 1.739 2 4.5 2s4.5-.847 4.5-2c0-.327-.18-.636-.5-.912"/><path stroke="currentColor" stroke-width="1.5" d="M3.321 6.658a4.833 4.833 0 0 1 3.918-4.197l.215-.04a24.736 24.736 0 0 1 9.091 0l.323.06a4.701 4.701 0 0 1 3.81 4.067c.418 3.544.43 7.125.034 10.672l-.017.154a4.838 4.838 0 0 1-4.215 4.26l-.906.113c-.495.062-.742.093-.99.118a24.88 24.88 0 0 1-5.169 0a51.167 51.167 0 0 1-.99-.118l-1.015-.126a4.714 4.714 0 0 1-4.105-4.137a46.932 46.932 0 0 1 0-10.689l.016-.137Z"/></g>`),
		g.Group(children),
	)
}