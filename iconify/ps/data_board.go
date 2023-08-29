package ps

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DataBoard(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M491 320h-24l-17-196q-1-25-19.5-42.5T388 64H277V21q0-21-21-21t-21 21v43H124q-24 0-42.5 17.5T62 124L45 320H21q-9 0-15 6t-6 15q0 10 6 16t15 6h79L64 484q-3 8 2 16t13 10h6q13 0 22-15l38-132h90v85q0 21 21 21t21-21v-85h90l38 134q3 15 22 15h6q19-7 15-26l-36-123h79q9 0 15-6t6-16q0-9-6-15t-15-6zm-404 0l18-194q0-8 6.5-13.5T126 107h262q9 0 15.5 5.5T410 126l17 194H87zm237-183q-49 65-53 70q-35-24-47-32l-15-13l-77 103q-5 6-3.5 15.5T137 294q7 5 12 5q12 0 17-9l52-68q38 26 47.5 31t16.5 3l8-2l4-7q3-4 13.5-17.5T333 195t23-31q5-6 4-15.5t-8-14.5q-5-6-14-5t-14 8z"/>`),
		g.Group(children),
	)
}