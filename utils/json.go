package utils

import (
	"bytes"
	"encoding/json"
)

func FastRaidJsonLoads(b []byte) map[string]interface{} {
	// Optimized parsing for our specific use case
	try := func() map[string]interface{} {
		if bytes.HasPrefix(b, []byte(`{"raid_id":`)) {
			parts := bytes.SplitN(b[11:], []byte(","), 2)
			if len(parts) > 0 {
				raidID := bytes.Trim(parts[0], `" }`)
				return map[string]interface{}{
					"raid_id": string(raidID),
				}
			}
		}

		// Standard parsing for other cases
		var result map[string]interface{}
		if err := json.Unmarshal(b, &result); err == nil {
			return result
		}
		return map[string]interface{}{}
	}

	// Catch any panics during parsing
	defer func() {
		recover()
	}()

	return try()
}
