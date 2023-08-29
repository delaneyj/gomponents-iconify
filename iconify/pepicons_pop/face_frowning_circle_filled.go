package pepicons_pop

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FaceFrowningCircleFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="none"><defs><mask id="pepiconsPopFaceFrowningCircleFilled0"><path fill="#fff" d="M0 0h26v26H0z"/><g fill="#000"><path fill-rule="evenodd" d="M13 20.5a7.5 7.5 0 1 0 0-15a7.5 7.5 0 0 0 0 15Zm0 2a9.5 9.5 0 1 0 0-19a9.5 9.5 0 0 0 0 19Z" clip-rule="evenodd"/><path d="M11.5 10.25a1.25 1.25 0 1 1-2.5 0a1.25 1.25 0 0 1 2.5 0Zm5.5 0a1.25 1.25 0 1 1-2.5 0a1.25 1.25 0 0 1 2.5 0Z"/><path fill-rule="evenodd" d="M15.221 17.627A1 1 0 0 0 16.8 16.4l-.8.6l.8-.601l-.002-.001l-.001-.003l-.004-.005l-.01-.011a1.891 1.891 0 0 0-.08-.095a2.538 2.538 0 0 0-.191-.188a3.597 3.597 0 0 0-.698-.478C15.192 15.286 14.28 15 13 15s-2.192.286-2.814.618a3.6 3.6 0 0 0-.698.478a2.557 2.557 0 0 0-.272.283l-.009.011l-.003.005l-.002.003H9.2c0 .001-.001.002.799.602l-.8-.6a1 1 0 0 0 1.579 1.227a1.598 1.598 0 0 1 .348-.245c.315-.168.902-.382 1.873-.382c.97 0 1.558.214 1.873.382a1.6 1.6 0 0 1 .348.245Z" clip-rule="evenodd"/></g></mask></defs><circle cx="13" cy="13" r="13" fill="currentColor" mask="url(#pepiconsPopFaceFrowningCircleFilled0)"/></g>`),
		g.Group(children),
	)
}