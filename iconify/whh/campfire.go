package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Campfire(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M998 756L59 1021q-16 4-32-10T4 973t-1-44t23-24l939-265q16-4 32 10t23 38t1 44t-23 24zM832 602q-1 11-4 27l-192 54q4-27 4-44q0-156-160-256q0 2-3 18t-5 25t-7 28t-10 32.5t-12 33t-16 34.5t-20 31.5t-25 30t-30 23.5t-26 35l-90-25q-1 0-1.5.5l-.5.5l1-1l-41-12q4-34 33-93.5t29-96.5q0-47-6.5-93.5t-15.5-79t-18.5-58.5t-16.5-39l-6-13q97 41 200 116q1-4 5.5-16t6-15.5T409 235t4-15t2-13t1-16q0-46-16-94t-32-73L353 0q10 4 28 12.5T450.5 50t97 61t99 82T735 295l1 1q24-31 64-57t68-37l27-11q-3 4-9.5 11t-23 30t-29.5 47t-24.5 60.5T795 411q40 105 37 191zm-598 48l-6 11q0-2 6-11zm-33 156L26 756q-17-4-23-24t1-44t23-38t32-10l364 103zm797 99q17 4 23 24t-1 44t-23 38t-32 10L585 914l222-63z"/>`),
		g.Group(children),
	)
}