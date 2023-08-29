package cil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Euro(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M316 120h60V88h-60c-73.747 0-136.794 46.657-161.195 112H104v32h42.292a172.176 172.176 0 0 0 0 56H104v32h50.805c24.4 65.343 87.448 112 161.2 112h60v-32H316a140.176 140.176 0 0 1-126.474-80H344v-32H178.815a140.661 140.661 0 0 1 0-56H344v-32H189.526A140.176 140.176 0 0 1 316 120Z"/>`),
		g.Group(children),
	)
}