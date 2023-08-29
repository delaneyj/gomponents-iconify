package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SmartSpeakerBroken(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" d="M3 14.05c.653.57 2.487 2.034 7 2.377M21 14c-.602.554-2.361 2.076-6.99 2.427"/><path d="M15.5 16c-.582-1.748-1.653-2.5-3.5-2.5s-2.918.752-3.5 2.5"/><path stroke-linecap="round" d="M8 4.412c-.32.275-.5.673-.5 1c0 1.153 1.739 2 4.5 2s4.5-.847 4.5-2c0-.327-.18-.636-.5-.912"/><path stroke-linecap="round" d="M21 12a46.819 46.819 0 0 1-.288 5.22l-.017.154a4.838 4.838 0 0 1-4.215 4.26l-.906.113c-.495.062-.742.093-.99.118a24.88 24.88 0 0 1-5.169 0a51.167 51.167 0 0 1-.99-.118l-1.015-.126a4.714 4.714 0 0 1-4.105-4.137a46.932 46.932 0 0 1 0-10.689l.016-.137a4.833 4.833 0 0 1 3.918-4.197l.215-.04a24.736 24.736 0 0 1 9.091 0l.323.06a4.701 4.701 0 0 1 3.81 4.067c.056.483.106.967.148 1.452"/></g>`),
		g.Group(children),
	)
}