/**
 * @flow
 */

/* eslint-disable */

'use strict';

/*::
import type { ConcreteRequest } from 'relay-runtime';
export type MeshType = "ALL" | "CITRIXSM" | "CONSUL" | "ISTIO" | "KUMA" | "LINKERD" | "NETWORKSM" | "NGINXSM" | "NONE" | "OCTARINE" | "OPENSERVICEMESH" | "TRAEFIK" | "%future added value";
export type Status = "DISABLED" | "ENABLED" | "UNKNOWN" | "%future added value";
export type AddonEventsSubscriptionVariables = {|
  selector?: ?MeshType
|};
export type AddonEventsSubscriptionResponse = {|
  +addonEvent: $ReadOnlyArray<{|
    +type: string,
    +status: ?Status,
    +config: {|
      +serviceName: string
    |},
  |}>
|};
export type AddonEventsSubscription = {|
  variables: AddonEventsSubscriptionVariables,
  response: AddonEventsSubscriptionResponse,
|};
*/


/*
subscription AddonEventsSubscription(
  $selector: MeshType
) {
  addonEvent: listenToAddonEvents(selector: $selector) {
    type
    status
    config {
      serviceName
    }
  }
}
*/

const node/*: ConcreteRequest*/ = (function(){
var v0 = [
  {
    "defaultValue": null,
    "kind": "LocalArgument",
    "name": "selector"
  }
],
v1 = [
  {
    "alias": "addonEvent",
    "args": [
      {
        "kind": "Variable",
        "name": "selector",
        "variableName": "selector"
      }
    ],
    "concreteType": "AddonList",
    "kind": "LinkedField",
    "name": "listenToAddonEvents",
    "plural": true,
    "selections": [
      {
        "alias": null,
        "args": null,
        "kind": "ScalarField",
        "name": "type",
        "storageKey": null
      },
      {
        "alias": null,
        "args": null,
        "kind": "ScalarField",
        "name": "status",
        "storageKey": null
      },
      {
        "alias": null,
        "args": null,
        "concreteType": "AddonConfig",
        "kind": "LinkedField",
        "name": "config",
        "plural": false,
        "selections": [
          {
            "alias": null,
            "args": null,
            "kind": "ScalarField",
            "name": "serviceName",
            "storageKey": null
          }
        ],
        "storageKey": null
      }
    ],
    "storageKey": null
  }
];
return {
  "fragment": {
    "argumentDefinitions": (v0/*: any*/),
    "kind": "Fragment",
    "metadata": null,
    "name": "AddonEventsSubscription",
    "selections": (v1/*: any*/),
    "type": "Subscription",
    "abstractKey": null
  },
  "kind": "Request",
  "operation": {
    "argumentDefinitions": (v0/*: any*/),
    "kind": "Operation",
    "name": "AddonEventsSubscription",
    "selections": (v1/*: any*/)
  },
  "params": {
    "cacheID": "3bd191d555f561b79be8da77b06e1a99",
    "id": null,
    "metadata": {},
    "name": "AddonEventsSubscription",
    "operationKind": "subscription",
    "text": "subscription AddonEventsSubscription(\n  $selector: MeshType\n) {\n  addonEvent: listenToAddonEvents(selector: $selector) {\n    type\n    status\n    config {\n      serviceName\n    }\n  }\n}\n"
  }
};
})();
// prettier-ignore
(node/*: any*/).hash = '1ecd4940edf0697c34d3f94a931ddb1d';

module.exports = node;
