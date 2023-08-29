package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Flashlight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSFlashlight0"><g fill="none" stroke-linejoin="round" stroke-width="4"><path fill="#fff" stroke="#fff" d="M17 22.196a.392.392 0 0 0-.157-.314A19.608 19.608 0 0 1 9 6.196V4h30v2.196a19.608 19.608 0 0 1-7.843 15.686a.392.392 0 0 0-.157.314V44H17V22.196Z"/><path stroke="#000" stroke-linecap="round" d="M38 11H10m14 17.008v4"/><path stroke="#fff" stroke-linecap="round" d="M17 22v0A20 20 0 0 1 9 6V4m30 0v2a20 20 0 0 1-8 16v0"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSFlashlight0)"/>`),
		g.Group(children),
	)
}