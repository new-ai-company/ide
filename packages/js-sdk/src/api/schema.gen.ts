/**
 * This file was auto-generated by openapi-typescript.
 * Do not make direct changes to the file.
 */


export interface paths {
  "/sandboxes": {
    /** @description List all running sandboxes */
    get: {
      parameters: {
        query?: {
          /** @description A query used to filter the sandboxes (e.g. "user=abc&app=prod"). Query and each key and values must be URL encoded. */
          query?: string;
        };
      };
      responses: {
        /** @description Successfully returned all running sandboxes */
        200: {
          content: {
            "application/json": components["schemas"]["RunningSandbox"][];
          };
        };
        400: components["responses"]["400"];
        401: components["responses"]["401"];
        500: components["responses"]["500"];
      };
    };
    /** @description Create a sandbox from the template */
    post: {
      requestBody: {
        content: {
          "application/json": components["schemas"]["NewSandbox"];
        };
      };
      responses: {
        /** @description The sandbox was created successfully */
        201: {
          content: {
            "application/json": components["schemas"]["Sandbox"];
          };
        };
        400: components["responses"]["400"];
        401: components["responses"]["401"];
        500: components["responses"]["500"];
      };
    };
  };
  "/sandboxes/{sandboxID}": {
    /** @description Get a sandbox by id */
    get: {
      parameters: {
        path: {
          sandboxID: components["parameters"]["sandboxID"];
        };
      };
      responses: {
        /** @description Successfully returned the sandbox */
        200: {
          content: {
            "application/json": components["schemas"]["RunningSandbox"];
          };
        };
        401: components["responses"]["401"];
        404: components["responses"]["404"];
        500: components["responses"]["500"];
      };
    };
    /** @description Kill a sandbox */
    delete: {
      parameters: {
        path: {
          sandboxID: components["parameters"]["sandboxID"];
        };
      };
      responses: {
        /** @description The sandbox was killed successfully */
        204: {
          content: never;
        };
        401: components["responses"]["401"];
        404: components["responses"]["404"];
        500: components["responses"]["500"];
      };
    };
  };
  "/sandboxes/{sandboxID}/logs": {
    /** @description Get sandbox logs */
    get: {
      parameters: {
        query?: {
          /** @description Starting timestamp of the logs that should be returned in milliseconds */
          start?: number;
          /** @description Maximum number of logs that should be returned */
          limit?: number;
        };
        path: {
          sandboxID: components["parameters"]["sandboxID"];
        };
      };
      responses: {
        /** @description Successfully returned the sandbox logs */
        200: {
          content: {
            "application/json": components["schemas"]["SandboxLogs"];
          };
        };
        401: components["responses"]["401"];
        404: components["responses"]["404"];
        500: components["responses"]["500"];
      };
    };
  };
  "/sandboxes/{sandboxID}/pause": {
    /** @description Pause the sandbox */
    post: {
      parameters: {
        path: {
          sandboxID: components["parameters"]["sandboxID"];
        };
      };
      responses: {
        /** @description The sandbox was paused successfully and can be resumed */
        204: {
          content: never;
        };
        401: components["responses"]["401"];
        404: components["responses"]["404"];
        409: components["responses"]["409"];
        500: components["responses"]["500"];
      };
    };
  };
  "/sandboxes/{sandboxID}/refreshes": {
    /** @description Refresh the sandbox extending its time to live */
    post: {
      parameters: {
        path: {
          sandboxID: components["parameters"]["sandboxID"];
        };
      };
      requestBody?: {
        content: {
          "application/json": {
            /** @description Duration for which the sandbox should be kept alive in seconds */
            duration?: number;
          };
        };
      };
      responses: {
        /** @description Successfully refreshed the sandbox */
        204: {
          content: never;
        };
        401: components["responses"]["401"];
        404: components["responses"]["404"];
      };
    };
  };
  "/sandboxes/{sandboxID}/resume": {
    /** @description Resume the sandbox */
    post: {
      parameters: {
        path: {
          sandboxID: components["parameters"]["sandboxID"];
        };
      };
      requestBody: {
        content: {
          "application/json": components["schemas"]["ResumedSandbox"];
        };
      };
      responses: {
        /** @description The sandbox was resumed successfully */
        201: {
          content: {
            "application/json": components["schemas"]["Sandbox"];
          };
        };
        401: components["responses"]["401"];
        404: components["responses"]["404"];
        409: components["responses"]["409"];
        500: components["responses"]["500"];
      };
    };
  };
  "/sandboxes/{sandboxID}/timeout": {
    /** @description Set the timeout for the sandbox. The sandbox will expire x seconds from the time of the request. Calling this method multiple times overwrites the TTL, each time using the current timestamp as the starting point to measure the timeout duration. */
    post: {
      parameters: {
        path: {
          sandboxID: components["parameters"]["sandboxID"];
        };
      };
      requestBody?: {
        content: {
          "application/json": {
            /**
             * Format: int32
             * @description Timeout in seconds from the current time after which the sandbox should expire
             */
            timeout: number;
          };
        };
      };
      responses: {
        /** @description Successfully set the sandbox timeout */
        204: {
          content: never;
        };
        401: components["responses"]["401"];
        404: components["responses"]["404"];
        500: components["responses"]["500"];
      };
    };
  };
  "/templates": {
    /** @description List all templates */
    get: {
      parameters: {
        query?: {
          teamID?: string;
        };
      };
      responses: {
        /** @description Successfully returned all templates */
        200: {
          content: {
            "application/json": components["schemas"]["Template"][];
          };
        };
        401: components["responses"]["401"];
        500: components["responses"]["500"];
      };
    };
    /** @description Create a new template */
    post: {
      requestBody: {
        content: {
          "application/json": components["schemas"]["TemplateBuildRequest"];
        };
      };
      responses: {
        /** @description The build was accepted */
        202: {
          content: {
            "application/json": components["schemas"]["Template"];
          };
        };
        401: components["responses"]["401"];
        500: components["responses"]["500"];
      };
    };
  };
  "/templates/{templateID}": {
    /** @description Rebuild an template */
    post: {
      parameters: {
        path: {
          templateID: components["parameters"]["templateID"];
        };
      };
      requestBody: {
        content: {
          "application/json": components["schemas"]["TemplateBuildRequest"];
        };
      };
      responses: {
        /** @description The build was accepted */
        202: {
          content: {
            "application/json": components["schemas"]["Template"];
          };
        };
        401: components["responses"]["401"];
        500: components["responses"]["500"];
      };
    };
    /** @description Delete a template */
    delete: {
      parameters: {
        path: {
          templateID: components["parameters"]["templateID"];
        };
      };
      responses: {
        /** @description The template was deleted successfully */
        204: {
          content: never;
        };
        401: components["responses"]["401"];
        500: components["responses"]["500"];
      };
    };
    /** @description Update template */
    patch: {
      parameters: {
        path: {
          templateID: components["parameters"]["templateID"];
        };
      };
      requestBody: {
        content: {
          "application/json": components["schemas"]["TemplateUpdateRequest"];
        };
      };
      responses: {
        /** @description The template was updated successfully */
        200: {
          content: never;
        };
        400: components["responses"]["400"];
        401: components["responses"]["401"];
        500: components["responses"]["500"];
      };
    };
  };
  "/templates/{templateID}/builds/{buildID}": {
    /** @description Start the build */
    post: {
      parameters: {
        path: {
          templateID: components["parameters"]["templateID"];
          buildID: components["parameters"]["buildID"];
        };
      };
      responses: {
        /** @description The build has started */
        202: {
          content: never;
        };
        401: components["responses"]["401"];
        500: components["responses"]["500"];
      };
    };
  };
  "/templates/{templateID}/builds/{buildID}/status": {
    /** @description Get template build info */
    get: {
      parameters: {
        query?: {
          /** @description Index of the starting build log that should be returned with the template */
          logsOffset?: number;
        };
        path: {
          templateID: components["parameters"]["templateID"];
          buildID: components["parameters"]["buildID"];
        };
      };
      responses: {
        /** @description Successfully returned the template */
        200: {
          content: {
            "application/json": components["schemas"]["TemplateBuild"];
          };
        };
        401: components["responses"]["401"];
        404: components["responses"]["404"];
        500: components["responses"]["500"];
      };
    };
  };
}

export type webhooks = Record<string, never>;

export interface components {
  schemas: {
    /**
     * Format: int32
     * @description CPU cores for the sandbox
     */
    CPUCount: number;
    EnvVars: {
      [key: string]: string;
    };
    Error: {
      /**
       * Format: int32
       * @description Error code
       */
      code: number;
      /** @description Error */
      message: string;
    };
    /**
     * Format: int32
     * @description Memory for the sandbox in MB
     */
    MemoryMB: number;
    NewSandbox: {
      envVars?: components["schemas"]["EnvVars"];
      metadata?: components["schemas"]["SandboxMetadata"];
      /** @description Identifier of the required template */
      templateID: string;
      /**
       * Format: int32
       * @description Time to live for the sandbox in seconds.
       * @default 15
       */
      timeout?: number;
    };
    Node: {
      /**
       * Format: int32
       * @description Number of allocated CPU cores
       */
      allocatedCPU: number;
      /**
       * Format: int32
       * @description Amount of allocated memory in MiB
       */
      allocatedMemoryMiB: number;
      /** @description Identifier of the node */
      nodeID: string;
      /**
       * Format: int32
       * @description Number of sandboxes running on the node
       */
      sandboxCount: number;
      status: components["schemas"]["NodeStatus"];
    };
    NodeDetail: {
      /** @description List of cached builds id on the node */
      cachedBuilds: string[];
      /** @description Identifier of the node */
      nodeID: string;
      /** @description List of sandboxes running on the node */
      sandboxes: components["schemas"]["RunningSandbox"][];
      status: components["schemas"]["NodeStatus"];
    };
    /**
     * @description Status of the node
     * @enum {string}
     */
    NodeStatus: "ready" | "draining";
    NodeStatusChange: {
      status: components["schemas"]["NodeStatus"];
    };
    ResumedSandbox: {
      /**
       * Format: int32
       * @description Time to live for the sandbox in seconds.
       * @default 15
       */
      timeout?: number;
    };
    RunningSandbox: {
      /** @description Alias of the template */
      alias?: string;
      /** @description Identifier of the client */
      clientID: string;
      cpuCount: components["schemas"]["CPUCount"];
      /**
       * Format: date-time
       * @description Time when the sandbox will expire
       */
      endAt: string;
      memoryMB: components["schemas"]["MemoryMB"];
      metadata?: components["schemas"]["SandboxMetadata"];
      /** @description Identifier of the sandbox */
      sandboxID: string;
      /**
       * Format: date-time
       * @description Time when the sandbox was started
       */
      startedAt: string;
      /** @description Identifier of the template from which is the sandbox created */
      templateID: string;
    };
    Sandbox: {
      /** @description Alias of the template */
      alias?: string;
      /** @description Identifier of the client */
      clientID: string;
      /** @description Version of the envd running in the sandbox */
      envdVersion: string;
      /** @description Identifier of the sandbox */
      sandboxID: string;
      /** @description Identifier of the template from which is the sandbox created */
      templateID: string;
    };
    /** @description Log entry with timestamp and line */
    SandboxLog: {
      /** @description Log line content */
      line: string;
      /**
       * Format: date-time
       * @description Timestamp of the log entry
       */
      timestamp: string;
    };
    SandboxLogs: {
      /** @description Logs of the sandbox */
      logs: components["schemas"]["SandboxLog"][];
    };
    SandboxMetadata: {
      [key: string]: string;
    };
    Team: {
      /** @description API key for the team */
      apiKey: string;
      /** @description Whether the team is the default team */
      isDefault: boolean;
      /** @description Name of the team */
      name: string;
      /** @description Identifier of the team */
      teamID: string;
    };
    TeamUser: {
      /** @description Email of the user */
      email: string;
      /**
       * Format: uuid
       * @description Identifier of the user
       */
      id: string;
    };
    Template: {
      /** @description Aliases of the template */
      aliases?: string[];
      /**
       * Format: int32
       * @description Number of times the template was built
       */
      buildCount: number;
      /** @description Identifier of the last successful build for given template */
      buildID: string;
      cpuCount: components["schemas"]["CPUCount"];
      /**
       * Format: date-time
       * @description Time when the template was created
       */
      createdAt: string;
      createdBy: components["schemas"]["TeamUser"] | null;
      /**
       * Format: date-time
       * @description Time when the template was last used
       */
      lastSpawnedAt: string;
      memoryMB: components["schemas"]["MemoryMB"];
      /** @description Whether the template is public or only accessible by the team */
      public: boolean;
      /**
       * Format: int64
       * @description Number of times the template was used
       */
      spawnCount: number;
      /** @description Identifier of the template */
      templateID: string;
      /**
       * Format: date-time
       * @description Time when the template was last updated
       */
      updatedAt: string;
    };
    TemplateBuild: {
      /** @description Identifier of the build */
      buildID: string;
      /**
       * @description Build logs
       * @default []
       */
      logs: string[];
      /**
       * @description Status of the template
       * @enum {string}
       */
      status: "building" | "ready" | "error";
      /** @description Identifier of the template */
      templateID: string;
    };
    TemplateBuildRequest: {
      /** @description Alias of the template */
      alias?: string;
      cpuCount?: components["schemas"]["CPUCount"];
      /** @description Dockerfile for the template */
      dockerfile: string;
      memoryMB?: components["schemas"]["MemoryMB"];
      /** @description Start command to execute in the template after the build */
      startCmd?: string;
      /** @description Identifier of the team */
      teamID?: string;
    };
    TemplateUpdateRequest: {
      /** @description Whether the template is public or only accessible by the team */
      public?: boolean;
    };
  };
  responses: {
    /** @description Bad request */
    400: {
      content: {
        "application/json": components["schemas"]["Error"];
      };
    };
    /** @description Authentication error */
    401: {
      content: {
        "application/json": components["schemas"]["Error"];
      };
    };
    /** @description Not found */
    404: {
      content: {
        "application/json": components["schemas"]["Error"];
      };
    };
    /** @description Conflict */
    409: {
      content: {
        "application/json": components["schemas"]["Error"];
      };
    };
    /** @description Server error */
    500: {
      content: {
        "application/json": components["schemas"]["Error"];
      };
    };
  };
  parameters: {
    buildID: string;
    nodeID: string;
    sandboxID: string;
    templateID: string;
  };
  requestBodies: never;
  headers: never;
  pathItems: never;
}

export type $defs = Record<string, never>;

export type external = Record<string, never>;

export type operations = Record<string, never>;
