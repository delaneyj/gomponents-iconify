package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bait(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M855 599q-57 58-126.5 97.5t-135.5 56T466 767t-107-24q-39 62-39 89q0 31 3 54t10.5 36.5t14 21.5t20.5 11.5t22.5 4t25.5.5q36 0 96-64q0 47-52.5 87.5T352 1024q-66 0-113-47t-47-113q0-17 5-37q-20 5-37 5q-66 0-113-47T0 672q0-55 40.5-107.5T128 512q-64 60-64 96q0 17 .5 25.5t4 22.5T80 676.5t21.5 14T138 701t54 3q27 0 89-40q-22-45-24-106t15-127t56.5-135.5T425 169q75-75 166-118.5T768 0v192q0 27 18.5 45.5T832 256h192q-7 86-50.5 177T855 599zM575.5 256q-26.5 0-45 18.5T512 320t18.5 45.5t45 18.5t45.5-18.5t19-45.5t-19-45.5t-45.5-18.5z"/>`),
		g.Group(children),
	)
}