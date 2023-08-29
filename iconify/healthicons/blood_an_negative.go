package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BloodANNegative(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><g clip-path="url(#healthiconsBloodANNegative0)"><path fill="currentColor" fill-rule="evenodd" d="M0 0h48v48H0V0Zm26.571 7.758a2 2 0 0 0 1.429.6h6c1.12 0 2 .896 2 1.97v17.21a8.441 8.441 0 0 0-.925-.625c-1.825-1.062-4.465-1.614-7.583.226c-2.568 1.515-4.983 1.925-7.61 1.98c-1.171.025-2.368-.02-3.651-.069l-.53-.02A68.676 68.676 0 0 0 12 28.972V10.327c0-1.073.88-1.97 2-1.97h6a2 2 0 0 0 1.429-.6l1.132-1.155a2.021 2.021 0 0 1 2.878 0l1.132 1.156ZM34 38h-4v2h-5v4h-2v-4h-5v-2h-4c-2.21 0-4-1.777-4-3.97V10.328c0-2.192 1.79-3.97 4-3.97h6l1.132-1.155a4.022 4.022 0 0 1 5.736 0L28 6.358h6c2.21 0 4 1.777 4 3.97V34.03c0 2.193-1.79 3.97-4 3.97ZM20 10a1 1 0 0 1 .923.615l3.742 8.98l.009.022l.008.02l1.241 2.978a1 1 0 0 1-1.846.77L23.083 21h-6.166l-.994 2.385a1 1 0 0 1-1.846-.77l1.241-2.978l.018-.042l3.74-8.98A1 1 0 0 1 20 10Zm2.25 9L20 13.6L17.75 19h4.5ZM33 16h-6a1 1 0 1 0 0 2h6a1 1 0 1 0 0-2Z" clip-rule="evenodd"/></g><defs><clipPath id="healthiconsBloodANNegative0"><path d="M0 0h48v48H0z"/></clipPath></defs></g>`),
		g.Group(children),
	)
}