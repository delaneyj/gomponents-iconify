package ps

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Delicious(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M444 142q-26-62-83-101q-30-19-60-29q-31-10-69-10q-48 0-90 18q-62 26-101 83q-19 31-29 61T2 232q0 48 18 90q26 62 83 101q31 19 61 29t68 10q48 0 90-18q62-26 101-83q19-30 29-60q10-31 10-69q0-48-18-90zm-28 168q-26 56-73 87q-23 16-52 25q-27 9-59 9V232H33q0-42 15-78q26-56 73-87q23-16 52-25q27-9 59-9v199h199q0 42-15 78z"/>`),
		g.Group(children),
	)
}