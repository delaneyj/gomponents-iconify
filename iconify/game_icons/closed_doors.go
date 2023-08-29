package game_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ClosedDoors(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M247 57.545c-29.212 2.622-71.312 17.137-106.37 38.172c-19.394 11.637-36.707 25.19-48.83 39.178C79.68 148.882 73 163 73 176v263h174v-78h-46v16h-66v-16h-32v-66h32v-16h66v16h46V57.545zm18 0V295h46v-16h66v16h32v66h-32v16h-66v-16h-46v78h174V176c0-13-6.678-27.118-18.8-41.105c-12.123-13.988-29.436-27.54-48.83-39.178c-35.06-21.035-77.16-35.55-106.37-38.172zM153 297v62h30v-62h-30zm176 0v62h30v-62h-30zm-208 16v30h14v-30h-14zm80 0v30h110v-30H201zm176 0v30h14v-30h-14zM73 457v30h174v-30H73zm192 0v30h174v-30H265z"/>`),
		g.Group(children),
	)
}