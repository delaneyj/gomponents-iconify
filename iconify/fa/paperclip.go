package fa

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Paperclip(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1536 1536"),
		g.Raw(`<path fill="currentColor" d="M1404 1257q0 117-79 196t-196 79q-135 0-235-100L117 656Q4 541 4 385q0-159 110-270T383 4q158 0 273 113l605 606q10 10 10 22q0 16-30.5 46.5T1194 822q-13 0-23-10L565 205q-79-77-181-77q-106 0-179 75t-73 181q0 105 76 181l776 777q63 63 145 63q64 0 106-42t42-106q0-82-63-145L633 531q-26-24-60-24q-29 0-48 19t-19 48q0 32 25 59l410 410q10 10 10 22q0 16-31 47t-47 31q-12 0-22-10L441 723q-63-61-63-149q0-82 57-139t139-57q88 0 149 63l581 581q100 98 100 235z"/>`),
		g.Group(children),
	)
}