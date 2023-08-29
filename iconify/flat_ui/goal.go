package flat_ui

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Goal(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 100 100"),
		g.Raw(`<defs><path id="flatUiGoal0" fill="#F0C419" fill-rule="evenodd" d="M89 74H76l13 13h13L89 74z" clip-rule="evenodd"/></defs><circle cx="50" cy="52" r="50" opacity=".15"/><circle cx="52" cy="50" r="50" fill="#3B97D3"/><path fill="#fff" d="M52 10c22.091 0 40 17.909 40 40S74.091 90 52 90S12 72.091 12 50s17.909-40 40-40z"/><path fill="#6BC8F2" fill-rule="evenodd" d="M52 20c16.569 0 30 13.431 30 30c0 16.568-13.431 30-30 30c-16.569 0-30-13.432-30-30c0-16.569 13.431-30 30-30z" clip-rule="evenodd"/><path fill="#fff" d="M52 30c11.046 0 20 8.955 20 20s-8.954 20-20 20s-20-8.955-20-20s8.954-20 20-20z"/><path fill="#E64C3C" fill-rule="evenodd" d="M52 40c5.523 0 10 4.477 10 10c0 5.522-4.477 10-10 10s-10-4.478-10-10c0-5.523 4.477-10 10-10z" clip-rule="evenodd"/><use href="#flatUiGoal0" fill-rule="evenodd" clip-rule="evenodd"/><path fill="#F29C1F" fill-rule="evenodd" d="M89 100V87L76 74v13l13 13z" clip-rule="evenodd"/><use href="#flatUiGoal0" fill-rule="evenodd" clip-rule="evenodd"/><path d="M50.439 48.439L15.121 83.757a50.513 50.513 0 0 0 2.097 2.145L52.56 50.56a1.5 1.5 0 1 0-2.121-2.121z" opacity=".15"/><path fill="#E57E25" d="M90.363 89.896c-.384 0-.769-.146-1.062-.439L50.939 51.061a1.501 1.501 0 0 1 2.123-2.121l38.363 38.395a1.501 1.501 0 0 1-1.062 2.561z"/>`),
		g.Group(children),
	)
}