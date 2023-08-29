package vs

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Cutlery(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1792 1792"),
		g.Raw(`<path fill="currentColor" d="M1280 64v1600q0 36-18 64q-16 30-46 46q-28 18-64 18t-64-18q-29-17-46-46q-18-28-18-64v-512H800q-13 0-22-10q-19-19 4-145q4-25 15.5-101.5t21.5-143T847 591t37-171.5t47.5-158t59.5-136t73-90T1152 0h64q26 0 45 19t19 45zm-640 0v640q0 38-8 65t-19 42.5t-32 27.5t-37 17t-46.5 13.5T448 885v779q0 36-18 64q-17 29-46 46q-29 18-64 18q-36 0-64-18q-29-17-46-46q-18-28-18-64V885q-19-7-49.5-15.5T96 856t-37-17t-32-27.5T8 769t-8-65V64q2-28 19-45T64 0q28 2 45 19q19 19 19 45v416q0 26 19 45t45 19t45-19t19-45V64q2-28 19-45t45-19q28 2 45 19t19 45v416q0 26 19 45t45 19t45-19t19-45V64q2-28 19-45t45-19q28 2 45 19t19 45z"/>`),
		g.Group(children),
	)
}