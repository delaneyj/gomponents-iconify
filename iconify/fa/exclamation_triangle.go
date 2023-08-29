package fa

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ExclamationTriangle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1536 1536"),
		g.Raw(`<path fill="currentColor" d="M1056 1375v-190q0-14-9.5-23.5t-22.5-9.5H832q-13 0-22.5 9.5T800 1185v190q0 14 9.5 23.5t22.5 9.5h192q13 0 22.5-9.5t9.5-23.5zm-2-374l18-459q0-12-10-19q-13-11-24-11H818q-11 0-24 11q-10 7-10 21l17 457q0 10 10 16.5t24 6.5h185q14 0 23.5-6.5t10.5-16.5zm-14-934l768 1408q35 63-2 126q-17 29-46.5 46t-63.5 17H160q-34 0-63.5-17T50 1601q-37-63-2-126L816 67q17-31 47-49t65-18t65 18t47 49z"/>`),
		g.Group(children),
	)
}