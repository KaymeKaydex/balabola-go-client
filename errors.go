package balabola

import (
	"fmt"
)

var BadQueryCompliance = map[int]error{
	0: nil,                // there is no error
	1: ErrSensitiveTopics, // balaboba does not accept requests for sensitive topics
}

var ErrSensitiveTopics = fmt.Errorf("Balaboba does not accept requests for sensitive topics," +
	" for example, about politics or religion. People may take the generated text too seriously." +
	"\n\nThe probability that a query sets one of the hot topics is determined by a neural " +
	"network trained on the estimates of random people. But she may overdo it or, conversely" +
	",miss something.")
