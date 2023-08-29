package ps

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Mixx(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M168 327q-16 0-34-9q24 6 51-4q32-13 43-38q0 21-17.5 36T168 327zm9 30q32 0 56.5-17.5T266 295q-24 44-79 52q-48 6-81-22q26 32 71 32zm115-105v9q0 47-33.5 81T178 376q-33 0-61-18.5T75 309q11 36 40.5 58.5T183 390q47 0 79.5-33t32.5-79q0-14-3-26zm66 37q0 73-51.5 125T182 466q-74 0-125.5-52T5 289q0-53 29-97t76-65L78 32q-3-8 1-15.5T91 7V6h6q15 0 20 15l22 97q21-6 43-6q73 0 124.5 52T358 289zm-34 0q0-59-41.5-101T182 146q-17 0-35 5l27 112q4 22-2 23q0 1-1 1q-7 0-12-18l-38-109q-37 18-59.5 52.5T39 289q0 59 42 101t101 42t100.5-42T324 289z"/>`),
		g.Group(children),
	)
}