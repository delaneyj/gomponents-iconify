package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MicrosoftOutlookLogoFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M88 144a16 16 0 1 1 16-16a16 16 0 0 1-16 16Zm144-24v88a16 16 0 0 1-5 11.61a8.8 8.8 0 0 1-.72.65A15.93 15.93 0 0 1 216 224H88a16 16 0 0 1-16-16v-16H40a16 16 0 0 1-16-16V80a16 16 0 0 1 16-16h56V48a16 16 0 0 1 16-16h80a16 16 0 0 1 16 16v56h8a15.94 15.94 0 0 1 10 3.54l.08.07l.14.11a6.78 6.78 0 0 1 .74.68A15.93 15.93 0 0 1 232 120ZM112 64h24a16 16 0 0 1 16 16v74.13l40-28.89V48h-80Zm-24 96a32 32 0 1 0-32-32a32 32 0 0 0 32 32Zm111.26 48L152 173.87V176a16 16 0 0 1-16 16H88v16ZM216 127.65L165.66 164L216 200.35Z"/>`),
		g.Group(children),
	)
}