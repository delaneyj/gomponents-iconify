package nonicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TemplateSixteen(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M1.547.309a1.64 1.64 0 0 0-.774.464a1.578 1.578 0 0 0-.457.747l-.063.2v12.56l.063.2c.1.322.208.498.457.747c.249.249.425.357.747.457l.2.063h12.56l.2-.063c.322-.1.498-.208.747-.457c.249-.249.357-.425.457-.747l.063-.2V1.72l-.063-.2a1.578 1.578 0 0 0-.457-.747a1.623 1.623 0 0 0-.786-.466C14.207.248 1.774.249 1.547.309M14.17 1.83l.07.071v12.198l-.067.067a.347.347 0 0 1-.15.084c-.108.022-11.938.022-12.046 0a.347.347 0 0 1-.15-.084l-.067-.067V1.901l.07-.071l.071-.07h12.198l.071.07"/>`),
		g.Group(children),
	)
}