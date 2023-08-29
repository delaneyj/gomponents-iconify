package ps

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MailStamp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M469 0H43Q25 0 12.5 12.5T0 43v298q0 18 12.5 30.5T43 384h426q18 0 30.5-12.5T512 341V43q0-18-12.5-30.5T469 0zM43 341V43h426v298H43zm341-149H256q-21 0-21 21q0 22 21 22h128q21 0 21-22q0-21-21-21zm-64 64h-64q-21 0-21 21q0 22 21 22h64q21 0 21-22q0-21-21-21zM427 64h-43q-21 0-21 21v64q0 22 21 22h43q9 0 15-6t6-16V85q0-9-6-15t-15-6z"/>`),
		g.Group(children),
	)
}