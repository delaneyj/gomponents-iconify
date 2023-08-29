package pepicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StudioLightSidePrint(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<g fill="currentColor"><g opacity=".8"><path fill-rule="evenodd" d="M14.816 5.864C16.136 4.894 18 5.837 18 7.476v2.048c0 1.64-1.863 2.583-3.184 1.612l-1.908-1.402l.907-1.234l-.907-1.234l1.908-1.402Zm-.21 2.636L16 9.524V7.476L14.606 8.5Z" clip-rule="evenodd"/><path d="m12.5 8.75l3.75-2.598v5.196L12.5 8.75Z"/><path fill-rule="evenodd" d="M15.414 18.908a1 1 0 0 1-1.32.507L11 18.04v.962a1 1 0 1 1-2 0v-.962l-3.094 1.375a1 1 0 1 1-.812-1.827L9 15.852v-3.85a1 1 0 0 1 2 0v3.85l3.906 1.736a1 1 0 0 1 .508 1.32Z" clip-rule="evenodd"/><path d="M7.5 7a1 1 0 0 1 1-1h3a1 1 0 0 1 1 1v3a1 1 0 0 1-1 1h-3a1 1 0 0 1-1-1V7Z"/><path fill-rule="evenodd" d="M6.5 7a2 2 0 0 1 2-2h3a2 2 0 0 1 2 2v3a2 2 0 0 1-2 2h-3a2 2 0 0 1-2-2V7Zm5 0h-3v3h3V7Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M12.293 6.709a1 1 0 0 1 0-1.415l2.693-2.693a1 1 0 1 1 1.415 1.414L13.707 6.71a1 1 0 0 1-1.414 0Zm3.914 7.498a1 1 0 0 1-1.414 0l-2.5-2.498a1 1 0 0 1 1.414-1.415l2.5 2.499a1 1 0 0 1 0 1.414Z" clip-rule="evenodd"/></g><path fill-rule="evenodd" d="M14.168 4.872c.997-.665 2.332.05 2.332 1.248v3.263c0 1.198-1.335 1.913-2.332 1.248l-2.945-1.963l.554-.832l2.946 1.963a.5.5 0 0 0 .777-.416V6.12a.5.5 0 0 0-.777-.416l-2.946 1.964l-.554-.832l2.945-1.964Zm-.711 13.083a.5.5 0 0 1-.66.253L9 16.521v1.73a.5.5 0 1 1-1 0v-1.73l-3.797 1.687a.5.5 0 1 1-.406-.913L8 15.427v-4.175a.5.5 0 0 1 1 0v4.175l4.203 1.868a.5.5 0 0 1 .254.66Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M5 5.752a1.5 1.5 0 0 1 1.5-1.5h4a1.5 1.5 0 0 1 1.5 1.5v4a1.5 1.5 0 0 1-1.5 1.5h-4a1.5 1.5 0 0 1-1.5-1.5v-4Zm1.5-.5a.5.5 0 0 0-.5.5v4a.5.5 0 0 0 .5.5h4a.5.5 0 0 0 .5-.5v-4a.5.5 0 0 0-.5-.5h-4Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M11.146 5.605a.5.5 0 0 1 0-.707l3-3a.5.5 0 0 1 .708.707l-3 3a.5.5 0 0 1-.708 0Zm3.708 8a.5.5 0 0 1-.708 0l-3-3a.5.5 0 0 1 .707-.707l3 3a.5.5 0 0 1 0 .707Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}