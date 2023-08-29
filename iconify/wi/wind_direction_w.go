package wi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WindDirectionW(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 30 30"),
		g.Raw(`<path d="M0 772q0-157 61.5-300t165-246.5T472.5 61T771 0t299 61t246 164.5T1480.5 472t61.5 300t-61.5 300t-165 245.5t-246 163.5t-298.5 61q-157 0-300-61t-246-164.5t-164-246T0 772zm170 0q0 245 179 424q177 177 422 177q163 0 302-81t220-219.5t81-300.5t-81-300.5T1073 252t-302-81q-162 0-300.5 81T251 471.5T170 772zm248 0q0-9 12-10l601-258q11-5 16.5 1.5t.5 16.5l-89 240q-7 10 0 20l89 238q5 10-.5 16t-16.5 1L430 782q-12-1-12-10z" fill="currentColor"/>`),
		g.Group(children),
	)
}