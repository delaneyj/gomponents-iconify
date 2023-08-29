package ps

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func World(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M256 0Q150 0 75 75T0 256t75 181t181 75t181-75t75-181t-75-181T256 0zM109 331q21 53 21 96q-40-28-63.5-73.5T43 256q0-103 81-166q18 14 33 36.5t14 42.5q-4 31-41 61q-19 20-25.5 47.5T109 331zm360-75q0 88-62.5 150.5T256 469q-44 0-83-17q5-65-26-136q-11-34 13-54q50-50 53-89q6-54-51-109q39-21 94-21q26 0 49 6l-43 43l-15 15l62 62l150 21q10 35 10 66zM331 130l-24-23l43-43q55 26 87 81zm-62 90l-17 25q-20 32-5 64l9 20l15 72q3 23 28 30q4 2 10 2q15 0 28-11l24-23l2-4q15-56 34-75q27-19 15-68q-7-27-37-39l-51-21h-4q-35 0-51 28zm19 49l17-26q4-8 13-8l43 15q8 4 8 8q2 7 2 24q-27 20-49 91l-11 11l-12-68l-13-26q-6-9 2-21z"/>`),
		g.Group(children),
	)
}