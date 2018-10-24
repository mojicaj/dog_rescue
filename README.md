# Dog Rescue API
API to manage dog records for a dog rescue

Action | HTTP Method | Endpoint
--- | --- | ---
Create Dog | POST | `/api/dog`
Get Dog | GET | `/api/dog/<Name>`
Update Dog | PUT | `/api/dog`
Remove Dog | DELETE | `/api/dog/<Name>`

The create and update endpoints require a json body payload formatted as follows:

```json
{
  "name": "Spot",
  "breed": "Beagle",
  "age": 1,
	"weight": "30 lbs" ,
	"condition": "Healthy",
	"description": "loyal",
	"status": "ready for adoption",
	"location": "Bronx, NY",
	"image_url": ""
}
```
