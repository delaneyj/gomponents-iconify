package il

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AddUser(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 750 850"),
		g.Raw(`<path fill="currentColor" d="M535 518q27 13 42 34t21 44t8 43t0 31t-8 16q-14 12-34 21t-44 16t-49 11l-49 8q-56 7-118 8q-63-1-119-8q-24-3-49-8t-48-11t-44-16t-35-21q-6-5-7.5-15.5T1 639t8-43t21-44t42-34l62-28q29-13 52-28t35-35t13-48v-6q-17-14-25-33l-23-54q-18-20-27-37q-8-14-8-26t17-12q-6-38-3-68q2-25 12-46t36-21q5-19 17-34q11-13 31-24t53-10q30 0 49 13t33 30t23 33t20 23q4 2 7 5t0 8q-4 14-8 41t1 50q18 1 17 12t-10 26q-10 17-28 37q-7 18-11 31t-9 24t-10 17t-15 15v6q0 29 13 48t36 35t51 28t62 28zM663 78q11 0 11 11v23q0 12-11 12h-58v58q0 5-3 8t-9 4h-23q-5 0-8-4t-4-8v-58h-57q-12 0-12-12V89q0-11 12-11h57V20q0-5 4-9t8-3h23q12 0 12 12v58h58z"/>`),
		g.Group(children),
	)
}