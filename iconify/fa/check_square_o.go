package fa

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CheckSquareO(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1536 1536"),
		g.Raw(`<path fill="currentColor" d="M1408 802v318q0 119-84.5 203.5T1120 1408H288q-119 0-203.5-84.5T0 1120V288Q0 169 84.5 84.5T288 0h832q63 0 117 25q15 7 18 23q3 17-9 29l-49 49q-10 10-23 10q-3 0-9-2q-23-6-45-6H288q-66 0-113 47t-47 113v832q0 66 47 113t113 47h832q66 0 113-47t47-113V866q0-13 9-22l64-64q10-10 23-10q6 0 12 3q20 8 20 29zm231-489l-814 814q-24 24-57 24t-57-24L281 697q-24-24-24-57t24-57l110-110q24-24 57-24t57 24l263 263l647-647q24-24 57-24t57 24l110 110q24 24 24 57t-24 57z"/>`),
		g.Group(children),
	)
}