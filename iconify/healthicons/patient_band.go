package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PatientBand(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><path d="M0 0h48v48H0z"/><path fill="currentColor" fill-rule="evenodd" d="M42 26.5a.5.5 0 0 1-.5.5H37v2h3.5a.5.5 0 0 1 .5.5v1a.5.5 0 0 1-.5.5h-8.447v2H37.3c-2.85 2.272-7.223 2.524-10.474.733l-5.845-3.22l-.965 1.752l5.845 3.22c3.947 2.174 9.276 1.888 12.806-1.019A16.04 16.04 0 0 0 40.221 33h.278a2.5 2.5 0 0 0 2.5-2.5v-1a2.49 2.49 0 0 0-.162-.888A2.498 2.498 0 0 0 44 26.5v-1c0-.889-.464-1.669-1.162-2.112A2.49 2.49 0 0 0 43 22.5v-1a2.5 2.5 0 0 0-2-2.45v-.81a3 3 0 0 0-2.767-2.992L35.039 15h-.653c-4.994 0-9.914 1.2-14.347 3.501l.922 1.775A29.143 29.143 0 0 1 34.386 17h.575l3.117.242a1 1 0 0 1 .922.997V19h-2v2h3.5a.5.5 0 0 1 .5.5v1a.5.5 0 0 1-.5.5H37v2h4.5a.5.5 0 0 1 .5.5v1ZM6 20h9v2H6v-2Zm0 10h9v2H6v-2Z" clip-rule="evenodd"/><path fill="currentColor" fill-rule="evenodd" d="M14 10a3 3 0 1 1 6 0v8a3 3 0 1 1-6 0v-8Zm4 .555a1 1 0 1 1-2 0a1 1 0 0 1 2 0ZM17 15a1 1 0 1 0 0-2a1 1 0 0 0 0 2Z" clip-rule="evenodd"/><path fill="currentColor" fill-rule="evenodd" d="M14 10a3 3 0 1 1 6 0v8a3 3 0 1 1-6 0v-8Zm4 6.83V18a1 1 0 1 1-2 0v-1.17a2.997 2.997 0 0 0 2 0Zm-.172-3.39a1 1 0 1 1-1.657 1.12a1 1 0 0 1 1.657-1.12Zm0-2.324a1 1 0 1 0-1.657-1.12a1 1 0 0 0 1.657 1.12Z" clip-rule="evenodd"/><path fill="currentColor" fill-rule="evenodd" d="M12 20a2 2 0 0 1 2-2h6a2 2 0 0 1 2 2v12a2 2 0 0 1-1.041 1.755c-.485.266-.959.693-.959 1.245a1 1 0 0 1-1 1h-4a1 1 0 0 1-1-1c0-.552-.475-.98-.959-1.245A2 2 0 0 1 12 32V20Zm3.5 1a.5.5 0 0 0-.5.5v9a.5.5 0 0 0 .5.5h3a.5.5 0 0 0 .5-.5v-9a.5.5 0 0 0-.5-.5h-3Z" clip-rule="evenodd"/><path fill="currentColor" fill-rule="evenodd" d="M15.236 32.986a3.3 3.3 0 0 1 .59 1.014h2.347a3.3 3.3 0 0 1 .591-1.014M14 18a2 2 0 0 0-2 2v12a2 2 0 0 0 1.041 1.755c.484.266.959.693.959 1.245a1 1 0 0 0 1 1h4a1 1 0 0 0 1-1c0-.552.474-.98.959-1.245A2 2 0 0 0 22 32V20a2 2 0 0 0-2-2h-6Zm1 3.5a.5.5 0 0 1 .5-.5h3a.5.5 0 0 1 .5.5v9a.5.5 0 0 1-.5.5h-3a.5.5 0 0 1-.5-.5v-9Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}