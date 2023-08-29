package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Alertpay(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="m1013 737l-174 23q-34 2-57-32q-24-34-57-83l-271 45l1 83q-5 49-37 56l-223 30q-14 3-21.5-10.5T172 812L366 44q6-23 15-34t20-9t23 11.5T451 40l559 640q10 12 13.5 25.5t0 22.5t-10.5 9zM508 329q-13-16-35.5-4.5T450 353q0 25 2 226l208-29Q521 345 508 329zM102 959l272-36q15-3 25-11v15q-6 54-41 61l-285 37q-34 7-56-23.5T3 942l218-797q10-36 27.5-55T288 74L69 891q-7 25 5.5 48.5T102 959zm599-128q27 40 66 37l192-32q-2 23-15.5 46.5T920 906l-195 25q-48 3-72-32q-18-34-46-84l-91 31q-1-33-1-57l127-19q14 21 29 36t22 20z"/>`),
		g.Group(children),
	)
}