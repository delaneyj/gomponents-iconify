package ion

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func IosToggle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path d="M128 320c-26.467 0-48 21.533-48 48s21.533 48 48 48 48-21.533 48-48-21.533-48-48-48z" fill="currentColor"/><path d="M383.5 272h-255C75.205 272 32 315.205 32 368.5S75.205 464 128.5 464h255c53.295 0 96.5-42.205 96.5-95.5S436.795 272 383.5 272zM128 432c-35.346 0-64-28.653-64-64 0-35.346 28.654-64 64-64s64 28.654 64 64c0 35.347-28.654 64-64 64z" fill="currentColor"/><path d="M384 192c26.467 0 48-21.533 48-48s-21.533-48-48-48-48 21.533-48 48 21.533 48 48 48z" fill="currentColor"/><path d="M128.5 240h255c53.295 0 96.5-42.205 96.5-95.5S436.795 48 383.5 48h-255C75.205 48 32 91.205 32 144.5S75.205 240 128.5 240zM384 80c35.346 0 64 28.654 64 64 0 35.347-28.654 64-64 64s-64-28.653-64-64c0-35.346 28.654-64 64-64z" fill="currentColor"/>`),
		g.Group(children),
	)
}