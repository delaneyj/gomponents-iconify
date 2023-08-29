package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Inbento(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M29.473 17.52c-.835-.417-3.756-.373-4.371-.395s-1.626.066-2.109.066a20.346 20.346 0 0 0-4.271.696L9.789 7.703l.794 19.636a36.352 36.352 0 0 0-1.12 6.633a3.743 3.743 0 0 0 .242 1.998a17.072 17.072 0 0 0 3.273 3.295c.944.55 8.83 1.032 12.322 1.032a77.832 77.832 0 0 0 9.159-.46a19.515 19.515 0 0 0 4.217-4.175a36.433 36.433 0 0 0-1.055-8.06l1.516-19.9Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M37.621 27.602a48.706 48.706 0 0 0-5.161-5.6l6.677-14.3M10.583 27.339s3.954-4.635 4.986-5.294L9.79 7.703m1.109 25.83L4.5 31.647m0 4.917l6.601-1.212M43.5 31.916l-6.264 1.617M43.5 36.8l-6.466-1.138m-13.021-2.248a8.405 8.405 0 0 0 1.528-.037M30.45 36.8c-3.856.994-5.473-1.179-5.673-1.633c-.777.791-3.235 2.45-5.222 1.89m-1.937-8.727l.017 2.492m13.134-2.307l-.05 2.206"/>`),
		g.Group(children),
	)
}