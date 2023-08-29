package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Twitter(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M1024 128q-91 0-121 4q58-35 89-132q-9 6-29.5 19t-33 20.5T898 55t-36 11Q800 0 709 0q-92 0-144.5 38.5T512 160q0 4-1 27t-1 42.5t2 26.5q-134-7-237-53T96 64q-32 55-32 96q0 57 16 94.5t56 64.5q-40 1-72 1q0 152 146 181q-27 8-56 8q-20 0-39-4q19 61 76.5 97.5T320 640q-23 23-54 37t-71 19.5t-65 6.5t-66 1q-4 0-29 1.5T0 706q27 52 110.5 89T322 832q112 0 207.5-36.5t162-97.5T805 558t69-164t22-170q0-3 9.5-5.5T930 211t31.5-13.5t34-26.5t28.5-43z"/>`),
		g.Group(children),
	)
}