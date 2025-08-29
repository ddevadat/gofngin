# Makefile for deploying Go OCI Function

# --- Configuration ---
CURRENT_CONTEXT := $(shell fn list context --output json | jq -r '.[] | select(.current==true) | .name')
COMPARTMENT_ID=$(shell yq eval -r '."oracle.compartment-id"' ~/.fn/contexts/$(CURRENT_CONTEXT).yaml | tr -d '\n')


APP_NAME := Demo-Fn-Application
FUNC_NAME := $(shell yq -r '.name' func.yaml)
REPO_NAME := $(shell echo $(APP_NAME) | tr '[:upper:]' '[:lower:]')/$(FUNC_NAME)

# --- Default target ---
.PHONY: deploy
deploy: ensure-repo
	@echo "Deploying function $(FUNC_NAME) to app $(APP_NAME)..."
	fn -v deploy --app $(APP_NAME)

# --- Ensure OCI Container Repository exists ---
.PHONY: ensure-repo
ensure-repo:
	@echo "Ensuring repository $(REPO_NAME) exists in compartment $(COMPARTMENT_ID)..."
	@oci artifacts container repository create \
		-c $(COMPARTMENT_ID) \
		--display-name $(REPO_NAME) \
		--output json 2>/dev/null || echo "Repository $(REPO_NAME) already exists. Skipping."


# --- Clean local build artifacts ---
.PHONY: clean
clean:
	@echo "Cleaning local build artifacts..."
	rm -rf ./bin ./tmp ./dist
	@echo "Clean complete."

# --- Build (optional) ---
.PHONY: build
build:
	@echo "Ensuring Go dependencies..."
	go mod tidy
	@echo "Building function binary..."
	go build -o ./bin/func ./func.go
	@echo "Build complete."
	
.PHONY: run-local
run-local: build
	@echo "Running function locally..."
	@./bin/func