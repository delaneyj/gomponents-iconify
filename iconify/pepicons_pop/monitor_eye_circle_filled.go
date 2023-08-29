package pepicons_pop

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MonitorEyeCircleFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="none"><defs><mask id="pepiconsPopMonitorEyeCircleFilled0"><path fill="#fff" d="M0 0h26v26H0z"/><g fill="#000"><path fill-rule="evenodd" d="M7.5 7.5A2.5 2.5 0 0 1 10 5h9.782a2.5 2.5 0 0 1 2.5 2.5v7.231a2.5 2.5 0 0 1-2.5 2.5H15.89v1.84a1 1 0 1 1-2 0v-1.846a1 1 0 0 1 .11-1.994h5.782a.5.5 0 0 0 .5-.5V7.5a.5.5 0 0 0-.5-.5H10a.5.5 0 0 0-.5.5v1.258a1 1 0 0 1-2 0V7.5Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M10.695 19.645a1 1 0 0 1 1-1h6.391a1 1 0 0 1 0 2h-6.39a1 1 0 0 1-1-1ZM6.138 15.36c.533.3 1.337.516 2.28.516c.942 0 1.746-.217 2.28-.517c.564-.317.637-.6.637-.686c0-.087-.073-.37-.638-.687c-.533-.3-1.337-.517-2.28-.517c-.942 0-1.746.217-2.28.517c-.564.318-.637.6-.637.687c0 .086.073.369.638.686Zm-.98 1.742c-.854-.48-1.658-1.3-1.658-2.43c0-1.13.804-1.95 1.657-2.43c.885-.497 2.04-.773 3.26-.773c1.221 0 2.376.276 3.26.774c.854.48 1.658 1.3 1.658 2.43c0 1.13-.804 1.95-1.657 2.43c-.885.497-2.04.773-3.26.773c-1.222 0-2.376-.276-3.26-.774Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M8.417 10.1a.9.9 0 0 1 .9.9v1.469a.9.9 0 0 1-1.8 0V11a.9.9 0 0 1 .9-.9Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M10.553 10.362a.9.9 0 0 1 .706 1.06l-.245 1.224a.9.9 0 0 1-1.765-.353l.245-1.225a.9.9 0 0 1 1.059-.706Zm-4.271 0a.9.9 0 0 0-.706 1.06l.245 1.224a.9.9 0 0 0 1.765-.353l-.245-1.225a.9.9 0 0 0-1.059-.706Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M4.037 11.208a.9.9 0 0 0-.309 1.234l.735 1.225a.9.9 0 0 0 1.543-.926l-.734-1.225a.9.9 0 0 0-1.235-.308Zm8.761 0a.9.9 0 0 1 .309 1.234l-.735 1.225a.9.9 0 0 1-1.543-.926l.734-1.225a.9.9 0 0 1 1.235-.308Z" clip-rule="evenodd"/><path d="M8.415 13.88a.75.75 0 1 1 0 1.5a.75.75 0 0 1 0-1.5Z"/><path fill-rule="evenodd" d="M8.665 14.63a.25.25 0 1 0-.5 0a.25.25 0 0 0 .5 0Zm-.25 1.75a1.75 1.75 0 1 1 0-3.5a1.75 1.75 0 0 1 0 3.5Z" clip-rule="evenodd"/></g></mask></defs><circle cx="13" cy="13" r="13" fill="currentColor" mask="url(#pepiconsPopMonitorEyeCircleFilled0)"/></g>`),
		g.Group(children),
	)
}