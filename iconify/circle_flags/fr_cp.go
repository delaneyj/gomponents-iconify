package circle_flags

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FrCp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<mask id="circleFlagsFrCp0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsFrCp0)"><path fill="#0052b4" d="M0 0h512v512H0Z"/><path fill="#fff" d="m256 52l204 204l-204 204L52 256Z"/><path fill="#ff9811" d="m232 463l12-299h24l12 299z"/><path fill="#6da544" d="M293 172c31-14 42-23 90-22c-40-18-77-23-108 0c9-23 27-54 59-77c-55 9-82 36-86 68c-14-32-55-45-100-45c50 27 36 27 73 63c-32-5-68 9-104 50c49-27 72-18 104-14a90 90 0 0 0-41 86c36-45 31-32 77-72c18 40 40 49 49 86c9-37 0-77-18-91c36 14 68 23 86 50c-9-63-50-72-81-82z"/></g>`),
		g.Group(children),
	)
}