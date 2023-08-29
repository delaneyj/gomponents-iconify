package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Pedometer(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.48 15.76c0-5.58-5.85-11.26-11.6-11.26A1.87 1.87 0 0 0 27 6.18l-1.52 15.87a15.16 15.16 0 0 1-5.26-4.33a7.83 7.83 0 0 1 .74-.58c.57-.39-1.08-3.78-1.77-4.61s-1.91.67-6.5 3.48s-5.69 5.8-4.94 6.64s3.61-.41 5.18-1.35s4.46-.07 4.73-.05c.08 0 .19-.15.33-.39c1.61 1.83 4.15 4.18 7.12 5l-1 10.59a8.16 8.16 0 0 1-1.11-.13c-.68-.13-1.9 3.44-2 4.51s1.77.88 7 2.17s7.84.09 8.18-1.2c.28-1-2.84-2.26-4.61-2.71s-3.21-3.1-3.39-3.31h-.3l.94-9.66c7.53-.75 11.66-5.68 11.66-10.36Zm-9.93-7.28c3.17.94 6.2 4.29 6.2 7.28c0 2.79-2.58 5.73-7.55 6.51Zm-10.33 9.24l-2.23 3.14m6.08 15.59l3.82-.71m-2.41-13.69l-.37 3.81m4.09-3.59l-.37 3.81"/>`),
		g.Group(children),
	)
}