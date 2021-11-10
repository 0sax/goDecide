# goDecide
Golang SDK for Indicina's Decide API

##How to use GoDecide

###Mono
1. Create a Mono Statement Object with NewMonoStatement()
2. Populate the Statement Object with transactions using the MonoStatement.AddTransaction() function
3. Create a Customer with NewCustomer(). You only need an Id.
4. Call the PaseMonoStatement Function
5. Handle the response

~~If you want to handle the json response from mono directly, use the ParseMonoString() Function.~~

###Custom Statement
1. Create a Custom Statement Object with NewCustomStatement()
2. Populate the Statement Object with transactions using the CustomStatement.AddTransaction() function
3. Create a Customer with NewCustomer(). You only need an Id.
4. Call the ParseCustomStatement Function
5. Handle the response

~~If you want to handle the json response from mono directly, use the ParseMonoString() Function.~~

##Example

    func monoHandler() httprouter.Handle {
        return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

		var yahyah []map[string]interface{}

		err := apiHelpers.Unmarshal(r.Body, &yahyah)
		if err != nil {
			apiHelpers.WriteError(w, err, 500, "internal error 1")
			return
		}

		cust := indicina.NewCustomer(
			"id1", "", "", "", "")

		ms := indicina.NewMonoStatement()

		for _, tx := range yahyah {
			ms.AddTransaction(
				int(tx["Amount"].(float64)),
				tx["ID"].(string),
				tx["Date"].(string),
				tx["Narration"].(string),
				tx["Type"].(string),
				tx["Category"].(string),
			)
		}

		ic, err := indicina.Login(
			os.Getenv("INDICINA_CLIENT_ID"),
			os.Getenv("INDICINA_CLIENT_SECRET"),
			os.Getenv("INDICINA_BASE_URL"),
		)
		if err != nil {
			apiHelpers.WriteError(w, err, 500, "internal error 2")
			return
		}

		sum, err := ic.ParseMonoStatement(cust, ms)
		if err != nil {
			apiHelpers.WriteError(w, err, 500, "internal error 3")
			return
		}

		apiHelpers.WriteOKJSONResponse(w, 200, "Mono Summary", sum)
		return

	    }
    }