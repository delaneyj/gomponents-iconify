package vs

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FaceDislike(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1792 1792"),
		g.Raw(`<path fill="currentColor" d="M832 0q169 0 323 66t265.5 177.5T1598 509t66 323t-66 323.5t-177.5 265.5t-265.5 177t-323 66t-323-66t-265.5-177T66 1155.5T0 832t66-323t177.5-265.5T509 66T832 0zM607 1265q122-111 225-113q102 0 225 113q20 19 44.5 14.5t38.5-29.5q30-50-12-89q-148-137-296-137t-296 137q-42 40-12 89q15 25 39 29.5t44-14.5zM384 704q0 53 37.5 90.5T512 832t90.5-37.5T640 704t-37.5-90.5T512 576t-90.5 37.5T384 704zm768 128q53 0 90.5-37.5T1280 704t-37.5-90.5T1152 576t-90.5 37.5T1024 704t37.5 90.5T1152 832z"/>`),
		g.Group(children),
	)
}