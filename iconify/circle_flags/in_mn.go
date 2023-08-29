package circle_flags

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InMn(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<mask id="circleFlagsInMn0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsInMn0)"><path fill="#d80027" d="M0 0h512v73L256 96L0 73Z"/><path fill="#eee" d="M0 73h512v73l-256 23L0 146Z"/><path fill="#333" d="M0 146h512v73l-256 23L0 219Z"/><path fill="#ffda44" d="M0 219h512v74l-256 22L0 293Z"/><path fill="#4a1f63" d="M0 293h512v73l-256 23L0 366Z"/><path fill="#338af3" d="M0 366h512v73l-256 23L0 439Z"/><path fill="#6da544" d="M0 439h512v73H0z"/></g>`),
		g.Group(children),
	)
}