package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Euro(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M832.5 384q-26.5 0-45.5-18.5T768 320q-21-83-94-137.5T512 128q-88 0-157 54t-91 138h184q27 0 45.5 19t18.5 45.5t-18.5 45T448 448H256v128h192q27 0 45.5 19t18.5 45.5t-18.5 45T448 704H264q22 84 91 138t157 54q89 0 162-54.5T768 704q0-26 19-45t45.5-19t45 19t18.5 45q0 23-15 41q-35 122-137 200.5T512 1024q-142 0-248.5-91.5T134 704H64q-27 0-45.5-18.5T0 640.5T18.5 595T64 576h64V448H64q-27 0-45.5-18.5T0 384.5T18.5 339T64 320h70q23-137 129.5-228.5T512 0q130 0 232 78.5T881 279q15 18 15 41q0 27-18.5 45.5t-45 18.5z"/>`),
		g.Group(children),
	)
}