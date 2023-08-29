package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Youcutvideoeditor(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="34.338" cy="13.118" r="7.618" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="34.338" cy="34.882" r="7.618" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="31.835" cy="31.509" r="1.415" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="34.338" cy="34.882" r=".75" fill="currentColor"/><circle cx="36.841" cy="31.4" r="1.415" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="38.474" cy="36.188" r="1.415" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="34.338" cy="39.127" r="1.415" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="30.203" cy="36.188" r="1.415" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="23.456" cy="24" r=".75" fill="currentColor"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M30.747 28.026L12.573 9.854H6.044L27.482 31.29m-4.026-4.025L12.573 38.147H6.044L20.191 24m3.265-3.265l4.026-4.026M26.72 24l4.027-4.027"/>`),
		g.Group(children),
	)
}