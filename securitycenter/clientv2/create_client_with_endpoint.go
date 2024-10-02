// Copyright 2024 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package clientv2

// [START securitycenter_set_client_endpoint_v2]
import (
	"context"
	"fmt"

	securitycenter "cloud.google.com/go/securitycenter/apiv2"
	"google.golang.org/api/option"
)

// createClientWithEndpoint creates a Security Command Center client for a
// regional endpoint, along with another client for the default endpoint.
func createClientWithEndpoint(repLocation string) error {
	// Instantiate client for default endpoint.
	ctx := context.Background()
	client, err := securitycenter.NewClient(ctx)
	if err != nil {
		return fmt.Errorf("securitycenter.NewClient: %w", err)
	}
	defer client.Close() // Closing the client safely cleans up background resources.

	// Assemble the regional endpoint URL using provided location.
	repEndpoint := fmt.Sprintf("securitycenter.%s.rep.googleapis.com:443", repLocation)
	// Instantiate client for regional endpoint.
	repCtx := context.Background()
	repClient, err := securitycenter.NewClient(repCtx, option.WithEndpoint(repEndpoint))
	if err != nil {
		return fmt.Errorf("securitycenter.NewClient: %w", err)
	}
	defer repClient.Close() // Closing the client safely cleans up background resources.

	return nil
}

// [END securitycenter_set_client_endpoint_v2]
