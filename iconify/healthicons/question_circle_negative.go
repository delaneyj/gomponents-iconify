package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func QuestionCircleNegative(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><g clip-path="url(#healthiconsQuestionCircleNegative0)"><path fill="currentColor" fill-rule="evenodd" d="M48 0H0v48h48V0Zm-4 24c0 11.046-8.954 20-20 20S4 35.046 4 24S12.954 4 24 4s20 8.954 20 20Zm-9-4.71c0-3.06-1.31-5.422-3.337-6.987c-1.96-1.515-4.48-2.207-6.917-2.294c-2.444-.086-4.99.428-7.103 1.537c-2.113 1.108-3.958 2.91-4.585 5.444a2 2 0 0 0 3.884.96c.273-1.107 1.128-2.112 2.559-2.862c1.43-.75 3.272-1.146 5.103-1.081c1.838.065 3.481.586 4.614 1.461C30.286 16.293 31 17.496 31 19.29c0 1.82-.946 3.14-2.41 4.118c-1.516 1.012-3.414 1.514-4.716 1.596A2 2 0 0 0 22 27v3a2 2 0 1 0 4 0v-1.27c1.552-.335 3.283-.975 4.81-1.995C33.046 25.244 35 22.811 35 19.29ZM24 34a2 2 0 1 0 0 4a2 2 0 0 0 0-4Z" clip-rule="evenodd"/></g><defs><clipPath id="healthiconsQuestionCircleNegative0"><path d="M0 0h48v48H0z"/></clipPath></defs></g>`),
		g.Group(children),
	)
}