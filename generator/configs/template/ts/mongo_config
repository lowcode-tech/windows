export const config = {
  port: 8080,
  secure: false,
  log: {
    level: 'info',
    map: {
      time: '@timestamp',
      msg: 'message'
    }
  },
  middleware: {
    log: true,
    skips: 'health,log',
    request: 'request',
    response: 'response',
    status: 'status',
    size: 'size'
  },
  mongo: {
    uri: 'mongodb://localhost:27017',
    db: 'masterdata'
  }
};

export const env = {
  sit: {
    mongo: {
      db: 'masterdata_sit',
    }
  },
  prd: {
    log: {
      level: 'error'
    },
    middleware: {
      log: false
    }
  }
};
